import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Dashboard } from './components/Dashboard';
import './index.css';

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <div className="min-h-screen bg-gray-50 py-12">
        <div className="max-w-4xl mx-auto px-4">
          <h1 className="text-4xl font-extrabold text-center text-gray-900 mb-12">URL Shortener</h1>
          <Dashboard />
        </div>
      </div>
    </QueryClientProvider>
  );
}

export default App;