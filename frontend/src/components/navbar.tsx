"use client";
import { useState } from "react";
import Link from "next/link";
import { Menu, X } from "lucide-react";
import Image from "next/image";

export default function Navbar() {
    const [isOpen, setIsOpen] = useState(false);

    return (
        <nav className="bg-bgGreen1 text-white p-4 border-b-2 border-white">
            <div className="container mx-auto flex justify-between items-center">

                <Link href="/" className="flex items-center gap-2 text-2xl font-bold font-epilogue">
                    <Image src="/Kale Logo.png" alt="Logo" width={18} height={17} priority />
                    <span>KaleCheck</span>
                </Link>

                <div className="hidden md:flex space-x-6">
                    <Link href="/" className="hover:text-gray-300 font-epilogue font-medium ">Home</Link>
                    <Link href="/learn" className="hover:text-gray-300 font-epilogue font-medium ">Learn</Link>
                    <Link href="/history" className="hover:text-gray-300 font-epilogue font-medium ">History</Link>
                    <Link href="/apidocs" className="hover:text-gray-300 font-epilogue font-medium">API Docs</Link>
                </div>


                <button className="md:hidden" onClick={() => setIsOpen(!isOpen)}>
                    {isOpen ? <X size={24} /> : <Menu size={24} />}
                </button>
            </div>


            {isOpen && (
                <div className="md:hidden flex flex-col space-y-4 mt-4 text-center">
                    <Link href="/" className="hover:text-gray-300">Home</Link>
                    <Link href="/about" className="hover:text-gray-300">About</Link>
                    <Link href="/contact" className="hover:text-gray-300">Contact</Link>
                </div>
            )}
        </nav>
    );
}
