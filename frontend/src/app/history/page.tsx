"use client"
import React, { useEffect, useState } from 'react';

const HistoryPage = () => {
  const [history, setHistory] = useState<any[]>([]);

  useEffect(() => {
    // Fetch upload history from your backend API
    const fetchHistory = async () => {
      try {
        const response = await fetch('http://localhost:8080/history');
        if (response.ok) {
          const data = await response.json();
          setHistory(data || []); // Ensure it's always an array
        }
      } catch (error) {
        console.error("Error fetching history:", error);
      }
    };

    fetchHistory();
  }, []);

  return (
    <div className="min-h-screen flex flex-col items-center py-10">
      <h1 className="text-2xl font-bold text-gray-800 mb-6">Upload History</h1>
      <div className="w-full max-w-2xl bg-white shadow-md rounded-lg p-6">
        {history && history.length === 0 ? (
          <p className="text-gray-500">No uploads yet.</p>
        ) : (
          <ul className="space-y-4">
            {history.map((item, index) => (
              <li key={index} className="flex items-center justify-between p-3 border-b">
                <span>{item.filename}</span>
                <span>{item.uploaded}</span>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
};

export default HistoryPage;

