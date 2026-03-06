"use client";
import dynamic from "next/dynamic";
import { useState } from "react";
import PointToPixel from "../../../util/pointToPixel";
import Sign from "@/components/sign";
import Navside from "@/components/navside";
import Navbar from "@/components/navbar";

const Viewer = dynamic(() => import("../../components/viewer"), {
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

  const sendFileToBE = async (filename: string, x: number, y: number) => {
    const res = await fetch(filename);
    if (!res.ok) throw new Error(`Gagal fetch file: ${res.status}`);
    const blob = await res.blob();

    const formData = new FormData();
    formData.append("file", blob, filename); 
    formData.append("x", x.toString());
    formData.append("y", "-"+y.toString());

    const response = await fetch("http://localhost:8080/sign", {
      method: "POST",
      body: formData,
    });

    if (!response.ok) throw new Error("Server error");

    const bl = await response.blob();
    const k = URL.createObjectURL(bl);
    window.open(k);
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
            <label htmlFor="">Point X = </label><input type="number" value={x} onChange={(e) => setX(parseInt(e.target.value))} />
            <label htmlFor="">Point Y = </label><input type="number" value={y} onChange={(e) => setY(parseInt(e.target.value))} />
            <label htmlFor="">Page = </label><input type="number" value={page} onChange={(e) => setPage(parseInt(e.target.value))} />
            <div className="p-5 border-t border-b border-(--line)">
              { fileUrl && (
                <button onClick={async () => (await sendFileToBE(fileUrl, x, y))} className="px-5 py-2 bg-(--secondary) rounded-xl border border-black">Ttd</button>
              )}
            </div>
          </div>
          <div className="w-[40%]">
            {fileUrl != null ? (
                <div className="w-full h-screen overflow-scroll flex justify-center items-center bg-gray-200 border-l border-(--line)">
                  <div className={`relative scale-75`}>
                    <Viewer path={fileUrl} pg={page}/>
                    <Sign x={PointToPixel(x)} y={PointToPixel(y)} />
                  </div>
                </div>
            ) : (
              <div className="w-full h-screen flex justify-center items-center bg-gray-200 border-l border-(--line)">
                <input type="file" accept="application/pdf" onChange={handleUpload} className="m-5 p-5 rounded-2xl border border-black border-dashed bg-(--secondary) font-bold text-black" />
              </div>
            )}
          </div>
        </div>
      </div>
    </section>
  );
}