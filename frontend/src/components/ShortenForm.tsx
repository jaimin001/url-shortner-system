import { useState } from 'react';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { shortenURL } from '../api';

export const ShortenForm = () => {
  const [url, setUrl] = useState('');
  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: shortenURL,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['links'] });
      setUrl('');
    },
  });

  return (
    <div className="p-4 bg-white rounded shadow-md">
      <h2 className="text-xl font-bold mb-4">Shorten your URL</h2>
      <input
        type="text"
        className="w-full p-2 border rounded mb-2"
        placeholder="Enter your long URL here"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
      />
      <button
        className="bg-blue-500 text-white px-4 py-2 rounded"
        onClick={() => mutation.mutate(url)}
        disabled={mutation.isPending}
      >
        {mutation.isPending ? 'Shortening...' : 'Shorten'}
      </button>
      {mutation.error && <p className="text-red-500 mt-2">Error shortening URL</p>}
      {mutation.isSuccess && <p className="text-green-500 mt-2">Successfully shortened!</p>}
    </div>
  );
};