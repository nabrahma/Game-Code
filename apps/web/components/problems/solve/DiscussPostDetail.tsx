import React, { useState, useEffect } from 'react';
import { ArrowLeft, MessageSquare, Send } from 'lucide-react';
import { UpvoteButton } from './UpvoteButton';
import { formatDistanceToNow } from 'date-fns';

interface Comment {
  id: string;
  user_id: string;
  content: string;
  upvotes: number;
  created_at: string;
  user?: { name: string; username: string };
}

interface Post {
  id: string;
  title: string;
  content: string;
  upvotes: number;
  created_at: string;
  user?: { name: string; username: string };
}

interface DiscussPostDetailProps {
  postId: string;
  onBack: () => void;
}

export function DiscussPostDetail({ postId, onBack }: DiscussPostDetailProps) {
  const [post, setPost] = useState<Post | null>(null);
  const [comments, setComments] = useState<Comment[]>([]);
  const [loading, setLoading] = useState(true);
  const [newComment, setNewComment] = useState('');
  const [submitting, setSubmitting] = useState(false);

  useEffect(() => {
    fetchData();
  }, [postId]);

  const fetchData = async () => {
    try {
      const [postRes, commentsRes] = await Promise.all([
        fetch(`http://localhost:8080/api/discuss/${postId}`),
        fetch(`http://localhost:8080/api/discuss/${postId}/comments`)
      ]);
      
      if (postRes.ok) setPost(await postRes.json());
      if (commentsRes.ok) setComments(await commentsRes.json());
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleComment = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!newComment.trim()) return;

    setSubmitting(true);
    try {
      const res = await fetch(`http://localhost:8080/api/discuss/${postId}/comments`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ content: newComment }),
      });
      if (res.ok) {
        setNewComment('');
        fetchData(); // Refresh comments
      }
    } catch (err) {
      console.error(err);
    } finally {
      setSubmitting(false);
    }
  };

  const handlePostUpvote = async () => {
    const res = await fetch(`http://localhost:8080/api/discuss/${postId}/upvote`, { method: 'POST' });
    const data = await res.json();
    return data.added;
  };

  const handleCommentUpvote = async (commentId: string) => {
    const res = await fetch(`http://localhost:8080/api/discuss/comments/${commentId}/upvote`, { method: 'POST' });
    const data = await res.json();
    return data.added;
  };

  if (loading) return <div className="p-4 text-zinc-500">Loading discussion...</div>;
  if (!post) return <div className="p-4 text-red-500">Post not found.</div>;

  return (
    <div className="flex flex-col h-full bg-[#0A0A0A]">
      <div className="sticky top-0 bg-zinc-900/90 backdrop-blur-sm border-b border-zinc-800 p-4 z-10">
        <button 
          onClick={onBack}
          className="flex items-center text-sm text-zinc-400 hover:text-white transition-colors group"
        >
          <ArrowLeft className="w-4 h-4 mr-1 group-hover:-translate-x-1 transition-transform" />
          Back to Discussions
        </button>
      </div>

      <div className="p-6">
        {/* Original Post */}
        <div className="mb-10">
          <h2 className="text-2xl font-bold text-white mb-2">{post.title}</h2>
          <div className="flex items-center gap-2 text-xs text-zinc-500 mb-6">
            <span className="font-medium text-emerald-400">@{post.user?.username || 'user'}</span>
            <span>•</span>
            <span>{formatDistanceToNow(new Date(post.created_at), { addSuffix: true })}</span>
          </div>
          
          <div className="text-zinc-300 leading-relaxed whitespace-pre-wrap mb-6 border border-zinc-800 bg-zinc-900/50 p-6 rounded-xl">
            {post.content}
          </div>
          
          <div className="flex items-center gap-4 border-b border-zinc-800 pb-8">
            <UpvoteButton initialUpvotes={post.upvotes} onToggle={handlePostUpvote} />
            <div className="flex items-center text-zinc-500 text-sm gap-1.5 font-medium">
              <MessageSquare className="w-4 h-4" />
              {comments.length} Comments
            </div>
          </div>
        </div>

        {/* Comments Section */}
        <div className="space-y-6 mb-8">
          <h3 className="text-lg font-semibold text-white">Responses</h3>
          {comments.length === 0 ? (
            <p className="text-zinc-500 text-sm">No comments yet. Be the first to reply!</p>
          ) : (
            comments.map(comment => (
              <div key={comment.id} className="bg-zinc-900 border border-zinc-800 rounded-lg p-5">
                <div className="flex justify-between items-start mb-3">
                  <div className="flex items-center gap-2 text-xs">
                    <span className="font-bold text-indigo-400">@{comment.user?.username || 'user'}</span>
                    <span className="text-zinc-600">•</span>
                    <span className="text-zinc-500">{formatDistanceToNow(new Date(comment.created_at))} ago</span>
                  </div>
                </div>
                <div className="text-zinc-300 text-sm whitespace-pre-wrap mb-4">
                  {comment.content}
                </div>
                <div>
                  <UpvoteButton initialUpvotes={comment.upvotes} onToggle={() => handleCommentUpvote(comment.id)} />
                </div>
              </div>
            ))
          )}
        </div>

        {/* Add Comment Form */}
        <form onSubmit={handleComment} className="mt-auto border-t border-zinc-800 pt-6">
          <div className="relative">
            <textarea
              placeholder="Write a reply..."
              className="w-full bg-zinc-900 border border-zinc-800 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-indigo-500 transition-colors h-24 resize-none pr-12"
              value={newComment}
              onChange={(e) => setNewComment(e.target.value)}
              required
            />
            <button
              type="submit"
              disabled={submitting || !newComment.trim()}
              className="absolute right-3 bottom-3 p-2 bg-indigo-600 hover:bg-indigo-700 disabled:bg-zinc-800 disabled:text-zinc-500 text-white rounded-lg transition-colors"
            >
              <Send className="w-4 h-4" />
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
