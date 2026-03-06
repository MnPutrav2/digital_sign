"use client";
import dynamic from "next/dynamic";
import { useState } from "react";
import PointToPixel from "../../util/pointToPixel";

const Viewer = dynamic(() => import("../components/viewer"), {
  ssr: false,
});

export default function Page() {
  const [fileUrl, setFileUrl] = useState<string | null>(null);

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


  const [x, setX] = useState<number>(100)
  const [y, setY] = useState<number>(100)

  function Sign({x, y}: {x: number, y: number}) {
    const sign = 90

    return (
      <>
      <div style={{
          position: "absolute",
          top: `${y}px`,
          left: `${x}px`,
          background: "rgba(255, 255, 255, 0.38)",
          border: "1px solid black",
          borderStyle: "dashed",
          width: `${sign}px`,
          height: `${sign}px`,
          margin:  "auto",
          color: "black",
          fontWeight: "bold",
          display: "flex",
          justifyContent: "center",
          alignItems: "center"
        }}>
          Here
        </div>
      </>
    )
  }

  return (
    <main>
      <div>
        {fileUrl != null ? (
          <>
            <label htmlFor="">Point X = </label><input type="number" value={x} onChange={(e) => setX(parseInt(e.target.value))} />
            <label htmlFor="">Point Y = </label><input type="number" value={y} onChange={(e) => setY(parseInt(e.target.value))} />
            <div className="w-[50%] h-[50%] overflow-scroll flex justify-center bg-gray-200">
              <div className="relative">
                <Viewer path={fileUrl} pg={1}/>
                <Sign x={PointToPixel(x)} y={PointToPixel(y)} />
              </div>
            </div>
          </>
        ) : (
          <div className="w-[50%] h-[50%] flex justify-center items-center bg-gray-200">
            <input type="file" accept="application/pdf" onChange={handleUpload} className="m-5 p-5 rounded-2xl border border-black border-dashed bg-red-100 font-bold text-black" />
          </div>
        )}
      </div>
    </main>
  );
}