"use client";


import React from "react";
import Image from "next/image";
import Link from "next/link";
import StepBox from "@/components/StepBox";
import { Upload, Search, CheckCircle, Lightbulb, Soup, FileUp, Check, ChartArea } from "lucide-react";
import Button from "@/components/Button";


const Homepage = () => {
  return (
    <div>
      <div className="flex flex-col items-center py-10 px-4">
        <div className="w-full max-w-screen-lg text-white flex flex-col items-center space-y-12">

          {/* Hero Image */}
          <div className="w-full flex justify-center">
            <Image
              src="/Depth 6, Frame 0.png"
              alt="Is your kale fresh?"
              width={928}
              height={480}
              priority
              className="max-w-full h-auto rounded-lg"
            />
          </div>

          {/* How it works */}
          <div className="w-full text-left">
            <h2 className="text-2xl font-bold font-epilogue pl-6">How it works</h2>
          </div>

          <div className="flex flex-wrap justify-center gap-6 w-full">
            <StepBox icon={<Search size={24} color="white" />} text="Upload your image" />
            <StepBox icon={<FileUp size={24} color="white" />} text="We'll analyze it" />
            <StepBox icon={<Check size={24} color="white" />} text="We'll tell you if it's fresh" />
            <StepBox icon={<Lightbulb size={24} color="white" />} text="Tips to keep it fresh longer" />
          </div>

          {/* Start Freshness Test Button */}
          <div>
            <Link href="/Upload-kale">
              <Button text="Start Freshness Test" />
            </Link>
          </div>

          {/* What is FreshKale? */}
          <div className="w-full text-left">
            <h2 className="text-2xl font-bold font-epilogue pl-6">What is FreshKale?</h2>
          </div>

          <div className="text-[16px] font-regular font-epilogue pl-6">
            <p>
              FreshKale is a fun, easy way to check the freshness of your kale. Just upload an image, and we'll analyze it for you.
              We’ll also give you tips on how to keep it fresh longer. Whether you're a seasoned chef or just starting out,
              FreshKale helps you make the most of your leafy greens.
            </p>
          </div>

          {/* Why use FreshKale? */}
          <div className="w-full text-left">
            <h2 className="text-2xl font-bold font-epilogue pl-6">Why use FreshKale?</h2>
          </div>

          <div className="flex flex-wrap justify-center gap-6 w-full">
            <StepBox icon={<ChartArea size={24} color="white" />} text="Keep your food fresh" />
            <StepBox icon={<Search size={24} color="white" />} text="Quickly find what you need" />
            <StepBox icon={<CheckCircle size={24} color="white" />} text="Create delicious dishes" />
            <StepBox icon={<Soup size={24} color="white" />} text="Enjoy your meals more" />
          </div>

          {/* Images */}
          <div className="flex flex-wrap justify-center gap-6 w-full">
            <div className="w-[30%] h-auto min-w-[150px]">
              <Image
                src="/Kale1.png"
                alt="Kale1"
                width={301}
                height={402}
                loading="lazy"
                className="w-full h-auto object-cover rounded-lg"
              />
            </div>
            <div className="w-[30%] h-auto min-w-[150px]">
              <Image
                src="/kale2.png"
                alt="Kale2"
                width={301}
                height={402}
                loading="lazy"
                className="w-full h-auto object-cover rounded-lg"
              />
            </div>
            <div className="w-[30%] h-auto min-w-[150px]">
              <Image
                src="/Kale3.png"
                alt="Kale3"
                width={301}
                height={402}
                loading="lazy"
                className="w-full h-auto object-cover rounded-lg"
              />
            </div>
          </div>

        </div>
      </div>
    </div>
  );
};

export default Homepage;
