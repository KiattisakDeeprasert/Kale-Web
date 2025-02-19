import { Input } from "@heroui/react";

export default function App() {
  const variants = ["underlined"];

  return (
    <div className="w-full flex flex-col gap-4">
      {variants.map((variant) => (
        <div key={variant} className="flex w-full flex-wrap md:flex-nowrap mb-6 md:mb-0 gap-4">

          <Input
            label="Weight"
            type="text"
            variant={variant}
            className="text-white !border-b-2 !border-stroke2 focus:!ring-0 focus:!border-stroke2"
          />

          <Input
            label="Temp"
            type="text"
            variant={variant}
            className="text-white !border-b-2 !border-stroke2 focus:!ring-0 focus:!border-stroke2"
          />
        </div>
      ))}
    </div>
  );
}
