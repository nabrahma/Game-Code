import React, { useState, useEffect } from 'react';
import { MessageSquare, ArrowUp, Plus } from 'lucide-react';
import { formatDistanceToNow } from 'date-fns';
import { DiscussPostForm } from './DiscussPostForm';
import { DiscussPostDetail } from './DiscussPostDetail';

interface Post {
  id: string;
  title: string;
  content: string;
  upvotes: number;
  created_at: string;
  user?: { name: string; username: string };
  comments?: any[];
}

export function DiscussFeed({ slug }: { slug: string }) {
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const [isComposing, setIsComposing] = useState(false);
  const [selectedPostId, setSelectedPostId] = useState<string | null>(null);

  useEffect(() => {
    fetchPosts();
  }, [slug]);

  const fetchPosts = async () => {
    setLoading(true);
    try {
      const res = await fetch(`http://localhost:8080/api/problems/${slug}/discuss`);
      if (res.ok) {
        const data = await res.json();
        setPosts(data || []);
      }
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  if (selectedPostId) {
    return <DiscussPostDetail postId={selectedPostId} onBack={() => setSelectedPostId(null)} />;
  }

  return (
    <div className="p-6 h-full flex flex-col">
      <div className="flex justify-between items-center mb-8">
        <div>
          <h2 className="text-xl font-bold text-white">Discussion</h2>
          <p className="text-sm text-zinc-400 mt-1">Share solutions, ask for help, or leave a hint.</p>
        </div>
        <button
          onClick={() => setIsComposing(!isComposing)}
          className="flex items-center gap-2 bg-zinc-800 hover:bg-zinc-700 text-white px-4 py-2 rounded-lg font-medium transition-colors border border-zinc-700 hover:border-zinc-600"
        >
          {isComposing ? 'Cancel' : (
            <>
              <Plus className="w-4 h-4" />
              New Post
            </>
          )}
        </button>
      </div>

      {isComposing && (
        <DiscussPostForm 
          slug={slug} 
          onSuccess={() => {
            setIsComposing(false);
            fetchPosts();
          }} 
        />
      )}

      {loading ? (
        <div className="text-zinc-500 text-center py-8">Loading discussions...</div>
      ) : posts.length === 0 ? (
        <div className="text-center py-12 border border-dashed border-zinc-800 rounded-xl bg-zinc-900/30">
          <MessageSquare className="w-8 h-8 text-zinc-600 mx-auto mb-3" />
          <h3 className="text-zinc-300 font-medium mb-1">No discussions yet</h3>
          <p className="text-zinc-500 text-sm">Be the first to start a conversation about this problem.</p>
        </div>
      ) : (
        <div className="space-y-4">
          {posts.map(post => (
            <div 
              key={post.id}
              onClick={() => setSelectedPostId(post.id)}
              className="bg-zinc-900 border border-zinc-800 hover:border-indigo-500/50 rounded-xl p-5 cursor-pointer transition-all group hover:shadow-[0_0_15px_rgba(99,102,241,0.1)]"
            >
              <h3 className="text-lg font-semibold text-zinc-100 group-hover:text-indigo-400 transition-colors mb-2">
                {post.title}
              </h3>
              <p className="text-zinc-400 text-sm line-clamp-2 mb-4 leading-relaxed">
                {post.content}
              </p>
              
              <div className="flex items-center justify-between text-xs">
                <div className="flex items-center gap-4 text-zinc-500">
                  <div className="flex items-center gap-1.5 bg-zinc-950 px-2.5 py-1 rounded-md border border-zinc-800">
                    <ArrowUp className="w-3.5 h-3.5" />
                    <span className="font-medium">{post.upvotes}</span>
                  </div>
                  <div className="flex items-center gap-1.5 font-medium">
                    <MessageSquare className="w-3.5 h-3.5" />
                    <span>{post.comments?.length || 0}</span>
                  </div>
                </div>
                
                <div className="flex items-center gap-2 text-zinc-500">
                  <span className="font-medium text-zinc-400">@{post.user?.username || 'user'}</span>
                  <span>•</span>
                  <span>{formatDistanceToNow(new Date(post.created_at))} ago</span>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
