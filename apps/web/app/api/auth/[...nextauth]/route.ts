import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";

const handler = NextAuth({
  providers: [
    CredentialsProvider({
      name: "Mock Dev Login",
      credentials: {
        username: { label: "Username", type: "text", placeholder: "testuser" },
      },
      async authorize(credentials) {
        if (!credentials?.username) return null;
        // Mock a user
        return {
          id: "00000000-0000-0000-0000-000000000001",
          name: credentials.username,
          email: `${credentials.username}@gamecode.dev`,
        };
      },
    }),
  ],
  session: { strategy: "jwt" },
  callbacks: {
    async session({ session, token }) {
      if (session.user) {
        session.user.id = token.sub as string;
      }
      return session;
    },
  },
  secret: "super-secret-dev-key",
});

export { handler as GET, handler as POST };
