"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import SubmitButton from "@/components/SubmitButton";
import ProgressLoader from "@/components/ProgressLoader";
import Input from "@/components/input";
import Button from "@/components/Button";

const UploadKale = () => {
  const [image, setImage] = useState<File | null>(null);
  const [isDragging, setIsDragging] = useState(false);
  const [isUploading, setIsUploading] = useState(false);
  const [imageUrl, setImageUrl] = useState<string | null>(null);
  const [isCompleted, setIsCompleted] = useState(false);
  const [uploadProgress, setUploadProgress] = useState(0); // Real progress
  const router = useRouter();

  const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (!files || files.length === 0) return alert("No file selected.");

    const file = files[0];
    if (!file.type.startsWith("image/")) return alert("Upload an image file.");
    if (file.size > 5 * 1024 * 1024) return alert("File size must be under 5MB.");

    setImage(file);
    const reader = new FileReader();
    reader.onloadend = () => {
      setImageUrl(reader.result as string);
      localStorage.setItem("imageUrl", reader.result as string);
    };
    reader.readAsDataURL(file);
  };

  const handleDrop = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    setIsDragging(false);
    if (event.dataTransfer.files.length > 0) {
      const file = event.dataTransfer.files[0];
      setImage(file);
      const reader = new FileReader();
      reader.onloadend = () => {
        setImageUrl(reader.result as string);
        localStorage.setItem("imageUrl", reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleSubmit = async () => {
    if (!image) return alert("Please upload an image first.");
    setIsUploading(true);
  
    const formData = new FormData();
    formData.append("file", image);
  
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/upload", true);
  

    xhr.upload.addEventListener("progress", (event) => {
      if (event.lengthComputable) {
        const percent = Math.round((event.loaded / event.total) * 100);
        setUploadProgress(percent); 
      }
    });
  
    xhr.onload = () => {
      if (xhr.status === 200) {
        const result = JSON.parse(xhr.responseText);
        console.log("Upload successful:", result);
        setIsCompleted(true);
        setTimeout(() => router.push("/result"), 3000);
      } else {
        console.error("Error uploading image:", xhr.statusText);
        alert("Upload failed");
      }
      setIsUploading(false);
    };
  
    xhr.send(formData);
  };

  return (
    <div className="bg-customGreen min-h-screen flex flex-col items-center pt-10 pb-5 px-4">
      <h1 className="text-lg md:text-2xl lg:text-3xl font-bold font-epilogue text-white mb-6 text-center">
        Evaluate Kale Freshness
      </h1>
      <span className="text-white text-center">
        Upload a photo of your kale to receive a freshness rating from 1-5.
      </span>

      <div
        className={`relative flex flex-col justify-center items-center w-full max-w-[928px] h-[200px] md:h-[250px] lg:h-[309px] border-2 rounded-lg p-4 transition
          ${isDragging ? "border-solid border-stroke2 bg-opacity-20 bg-white" : "border-dashed border-stroke2"}`}
        onDragOver={(e) => { e.preventDefault(); setIsDragging(true); }}
        onDragLeave={() => setIsDragging(false)}
        onDrop={handleDrop}
        role="button"
        aria-label="Upload area"
        tabIndex={0}
      >
        {imageUrl ? (
          <img src={imageUrl} alt="Preview" className="w-full h-full object-contain rounded-md" />
        ) : (
          <p className="text-white text-center font-epilogue font-bold text-sm md:text-base">
            Upload an image of your kale (png, jpg, jpeg)
          </p>
        )}

        <input type="file" accept="image/*" onChange={handleImageChange} className="absolute inset-0 w-full h-full opacity-0 cursor-pointer" />
      </div>

      <div className="mt-4 w-full max-w-[400px]">
        <Input label="Your Name" placeholder="Enter your name" type="text" />
      </div>

      <ProgressLoader isUploading={isUploading} progress={uploadProgress} onComplete={() => {}} /> 

      {isCompleted && (
        <div className="mt-6 text-white text-xl font-bold blinking font-epilogue flex items-center justify-center text-center">
          Upload completed! Hold a second for your result.
        </div>
      )}

      {!isUploading && (
        <div className="mt-10">
          <Button text="Submit" onClick={handleSubmit} />
        </div>
      )}
    </div>
  );
};

export default UploadKale;
