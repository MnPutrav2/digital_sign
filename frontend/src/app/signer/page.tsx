"use client";
import dynamic from "next/dynamic";
import { useRef, useState } from "react";
import PointToPixel from "../../../util/pointToPixel";
import Sign from "@/components/Sign/sign";
import Navside from "@/components/Navside/navside";
import Navbar from "@/components/Navbar/navbar";

const Viewer = dynamic(() => import("./viewer"), {
  ssr: false,
});

export default function Signer() {
  const [fileUrl, setFileUrl] = useState<string | null>(null);
  const [page, setPage] = useState<number>(1)
  const [x, setX] = useState<number>(100)
  const [y, setY] = useState<number>(100)

  const handleUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    const formData = new FormData();
    formData.append("file", file);

    const res = await fetch("/api/upload", {
      method: "POST",
      body: formData,
    });

    const data = await res.json();
    if (data.url) setFileUrl(data.url);
  };

  return (
    <section className="flex">
      <Navside/>
      <div className="w-full">
        <div>
          <Navbar/>
        </div>
        <div className="flex w-full">
          <div className="w-[60%]">
          </div>
          <div className="w-[40%]">
            {fileUrl != null ? (
              <Viewer path={fileUrl}/>
            ) : (
              <div className="w-full h-screen flex justify-center items-center bg-gray-200 border-l border-(--line)">
                <input type="file" accept="application/pdf" onChange={handleUpload} className="m-5 p-5 rounded-2xl border border-(--line) border-dashed bg-(--secondary) font-bold text-black" />
              </div>
            )}
          </div>
        </div>
      </div>
    </section>
  );
}