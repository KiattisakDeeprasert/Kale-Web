"use client";

import { Card, CardFooter } from "@heroui/react";

interface BenefitCardProps {
  title: string;
  subtitle: string;
  imageUrl: string;
}

const BenefitCard: React.FC<BenefitCardProps> = ({
  title,
  subtitle,
  imageUrl,
}) => {
  return (
    <Card
      isFooterBlurred
      className="relative border-none overflow-hidden shadow-lg h-56 w-100 flex items-end 
        transition-transform duration-300 ease-in-out hover:scale-105 hover:shadow-2xl"
      style={{
        backgroundImage: `url(${imageUrl})`,
        backgroundSize: "cover",
        backgroundPosition: "center",
      }}
    >
      <CardFooter
        className="absolute bottom-0 w-full bg-black/40 text-white p-4 flex flex-col items-start 
        transition-all duration-300 ease-in-out hover:bg-black/60"
      >
        <p className="text-sm uppercase opacity-80">{subtitle}</p>
        <h3 className="text-lg font-semibold">{title}</h3>
      </CardFooter>
    </Card>
  );
};

export default BenefitCard;
