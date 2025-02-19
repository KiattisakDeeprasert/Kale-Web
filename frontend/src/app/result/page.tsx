"use client";

import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { HoverEffect } from "@/components/card-hover-effect";
import Button from "@/components/Button";


const fetchFreshnessData = async () => {
  const response = await fetch("https://catfact.ninja/fact");
  const data = await response.json();
  return data;
};

const Result = () => {
  const [freshnessRating, setFreshnessRating] = useState<number | null>(null);
  const [imageUrl, setImageUrl] = useState<string | null>(null);
  
  const [items, setItems] = useState<{ title: string; description: string; link: string }[]>([]); 
  const router = useRouter();
  const [isDragging, setIsDragging] = useState(false);
  const [image, setImage] = useState<File | null>(null);
  const [isUploading, setIsUploading] = useState(false);
  const [isCompleted, setIsCompleted] = useState(false);

  useEffect(() => {
    setFreshnessRating(Math.floor(Math.random() * 5) + 1);
    setImageUrl(localStorage.getItem("imageUrl") || null);

    const fetchData = async () => {
      const data = await fetchFreshnessData();
      const formattedData = [
        {
          title: "Date of purchase",
          description: data.fact || "No information available",
          link: "/purchase",
        },
        {
          title: "Expiration date",
          description: data.fact || "No information available",
          link: "/expiration",
        },
        {
          title: "Estimated freshness",
          description: data.fact || "No information available",
          link: "/freshness",
        },
      ];
      setItems(formattedData);
    };

    fetchData();
  }, []);

  const handleDragOver = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    setIsDragging(true);
  };

  const handleSubmit = () => {
    if (image) {
      setIsUploading(true);

      setIsUploading(false);
      setIsCompleted(true);

      router.push("/Upload-kale");
    }
  };

  const handleGoBack = () => {
    router.push("/Upload-kale");
  };

  return (
    <div className="bg-customGreen min-h-screen flex flex-col items-center pt-10 pb-5 px-4">
      <h1 className="text-lg md:text-2xl lg:text-3xl font-bold font-epilogue text-white mb-6 text-center">
        Kale Freshness Result
      </h1>
  
      <div
        className={`relative flex flex-col justify-center items-center w-full max-w-[928px] h-[200px] md:h-[250px] lg:h-[309px] border-2 rounded-lg p-4 transition
          ${
            isDragging
              ? "border-solid border-stroke2 bg-opacity-20 bg-white"
              : "border-dashed border-stroke2"
          }`}
      >
        {imageUrl ? (
          <img
            src={imageUrl}
            alt="Preview"
            className="w-full h-full object-contain rounded-md"
          />
        ) : (
          <p className="text-white text-center font-epilogue font-bold text-sm md:text-base">
            Upload an image of your kale (png, jpg, jpeg)
          </p>
        )}
      </div>
  
      {freshnessRating ? (
        <div className="mt-6 text-white text-xl font-bold">
          Your kale freshness rating is: {freshnessRating} / 5
        </div>
      ) : (
        <div className="mt-6 text-white text-xl font-bold">
          Loading freshness rating...
        </div>
      )}
  
      <div className="mt-6 w-full max-w-[928px]">
        <HoverEffect items={items} className="h-full w-full" />
      </div>
  
      {/* Fixed button positioning */}
      <div className="mt-1 z-5">
        <Button
          text="Go back to Upload"
          onClick={handleGoBack}
        />
      </div>
    </div>
  );
}

export default Result;
