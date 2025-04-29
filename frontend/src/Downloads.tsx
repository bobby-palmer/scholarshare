import React, { useState } from 'react';

interface PdfSearchResult {
  id: string;
  title: string;
  description: string;
  downloadUrl: string;
}

const DownloadPage: React.FC = () => {
  const [query, setQuery] = useState<string>('');
  const [results, setResults] = useState<PdfSearchResult[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const handleSearch = async () => {
    if (query.trim() === '') return;

    setLoading(true);

    try {
      const response = await fetch(`http://localhost:8080/api/search-pdfs?query=${encodeURIComponent(query)}`);
      if (response.ok) {
        const data: PdfSearchResult[] = await response.json();
        setResults(data);
      } else {
        console.error('Failed to fetch PDFs');
      }
    } catch (error) {
      console.error('Error during search:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <div>
        <input
          type="text"
          placeholder="Search PDFs..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
        />
        <button onClick={handleSearch} disabled={loading}>
          {loading ? 'Searching...' : 'Search'}
        </button>
      </div>

      <div>
        {results.length > 0 ? (
          <ul>
            {results.map((pdf) => (
              <li key={pdf.id}>
                <a href={pdf.downloadUrl} target="_blank" rel="noopener noreferrer">
                  {pdf.title}
                </a>
                <p>{pdf.description}</p>
              </li>
            ))}
          </ul>
        ) : (
          <p>No results found.</p>
        )}
      </div>
    </div>
  );
};

export default DownloadPage;

