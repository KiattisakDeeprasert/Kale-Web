import React, { useEffect, useState } from "react";
import { Progress } from "@heroui/react";

interface ProgressLoaderProps {
  isUploading: boolean;
  progress: number;
  onComplete: () => void;
}

const ProgressLoader: React.FC<ProgressLoaderProps> = ({ isUploading, progress, onComplete }) => {
  return (
    <div className="w-full max-w-[928px] mt-4">
      {isUploading && (
        <Progress
          aria-label="Uploading..."
          className="w-full"
          color="success"
          showValueLabel={true}
          size="md"
          value={progress} // Use real progress here
        />
      )}
    </div>
  );
};

export default ProgressLoader;
