"use client";

export default function Banner() {
  return (
    <div
      className="w-full h-80 md:h-96 flex items-center justify-center bg-cover bg-center relative"
      style={{ backgroundImage: "url('https://picsum.photos/800/400')" }}
    >
      {/* Overlay for better text contrast */}
      <div className="absolute inset-0 bg-black/40 backdrop-blur-xl"></div>

      {/* Content Box */}
      <div className="relative z-10 text-center p-6 rounded-lg max-w-lg">
        <h1 className="text-3xl md:text-4xl font-bold">Is your kale fresh?</h1>
        <p className="mt-2 text-lg text-gray-300">
          Upload an image and we'll analyze its freshness.
        </p>
      </div>
    </div>
  );
}
