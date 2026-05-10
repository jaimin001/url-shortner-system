import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { shortenURL, fetchRecentLinks } from '../api';

export const Dashboard = () => {
  const [url, setUrl] = useState('');
  const queryClient = useQueryClient();

  const { data: links = [] } = useQuery({
    queryKey: ['links'],
    queryFn: fetchRecentLinks,
  });

  const mutation = useMutation({
    mutationFn: (url: string) => shortenURL(url),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['links'] });
      setUrl('');
    },
    onError: (err: any) => {
      console.error("Mutation error:", err);
      alert(err.response?.data?.error || "Failed to shorten URL");
    }
  });

  return (
    <div className="space-y-8">
      <div className="p-6 bg-white rounded-xl shadow-sm border border-gray-100">
        <h2 className="text-xl font-semibold mb-4">Create New Link</h2>
        <div className="flex gap-2">
          <input
            className="flex-1 p-2 border rounded-lg"
            placeholder="Enter long URL"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
          />
          <button
            className="bg-indigo-600 text-white px-6 py-2 rounded-lg font-medium hover:bg-indigo-700"
            onClick={() => mutation.mutate(url)}
          >
            Shorten
          </button>
        </div>
      </div>

      <div className="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
        <h2 className="text-xl font-semibold p-6 border-b">Recent Links</h2>
        <table className="w-full">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left">Short URL</th>
              <th className="px-6 py-3 text-left">Original</th>
            </tr>
          </thead>
          <tbody>
            {links.map((link: any, i: number) => (
              <tr key={i} className="border-t">
                <td className="px-6 py-3 text-indigo-600 font-mono">
                  <a href={`http://localhost:8080/${link.ID}`} target="_blank" rel="noreferrer">
                    {link.ID}
                  </a>
                </td>
                <td className="px-6 py-3 truncate max-w-xs" title={link.Original}>
                  {link.Original}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};