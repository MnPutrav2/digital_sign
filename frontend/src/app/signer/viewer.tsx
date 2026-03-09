"use client";

import { Document, Page, pdfjs } from "react-pdf";
import "react-pdf/dist/Page/TextLayer.css";
import "react-pdf/dist/Page/AnnotationLayer.css";
import { useRef, useState } from "react";
import PointToPixel from "../../../util/pointToPixel";
import Sign from "../../components/Sign/sign";
import PixelToPoint from "../../../util/pixelToPoint";
import Modal from "@/components/Modal/primary";
import { useClickOutside } from "../../../hooks/useClickOutside";

pdfjs.GlobalWorkerOptions.workerSrc = new URL("pdfjs-dist/build/pdf.worker.min.mjs",import.meta.url).toString();

export default function Viewer({ path }: { path: string }) {
  const [width, setWidth] = useState<number>(100)
  const [height, setHeight] = useState<number>(100)
  const [numPage, setNumPage] = useState<number>(1)
  const [modal, setModal] = useState<boolean>(false)
  const [x, setX] = useState<number>(100)
  const [y, setY] = useState<number>(100)
  const [page, setPage] = useState<number>(1)
  const [barcode, setBarcode] = useState<boolean>(false)
  const [passphrase, setPassphrase] = useState<string>("")
  const pageRef = useRef<HTMLDivElement>(null);
  const modalRef = useRef<HTMLDivElement>(null);
  useClickOutside(modalRef, () => setModal(false))

  const handleClick = (e: React.MouseEvent<HTMLDivElement>) => {
    const rect = pageRef.current?.getBoundingClientRect();
    if (!rect) return;

    const wx = e.clientX - rect.left;
    const wy = e.clientY - rect.top;

    setX(wx)
    setY(wy)
  };

  async function SignDocument() {

    if(passphrase == ""){
      alert("passphrase harus diisi")
      return
    }

    const res = await fetch(path);
    if (!res.ok) throw new Error(`Gagal fetch file: ${res.status}`);
    const blob = await res.blob();

    const formData = new FormData();
    formData.append("file", blob, path); 
    formData.append("x", x.toString());
    formData.append("y", "-"+y.toString());
    formData.append("page", page.toString());

    const response = await fetch("http://localhost:8080/sign", {
      method: "POST",
      body: formData,
    });

    if (!response.ok) throw new Error("Server error");

    setModal(false)
    const bl = await response.blob();
    const k = URL.createObjectURL(bl);
    window.open(k);
  };

  const onLoadSuccess = async (pdf: pdfjs.PDFDocumentProxy) => {
    const p = await pdf.getPage(page)
    const viewport = p.getViewport({ scale: 1 })
    setNumPage(pdf.numPages)

    setWidth(viewport.width)
    setHeight(viewport.height)
  }

  function documentPage(p: number) {
    if(p == 0){
      alert("out of range")
      setPage(1)
      return
    }
    
    if(p > numPage){
      alert("out of range")
      setPage(1)
      return
    }

    setPage(p)
  }

  return (
    <section className="h-screen">
      <div className="flex bg-white p-2 border-b border-l border-(--line)">
        <div className="m-1 mx-2"><input className="w-10" type="number" value={page} onChange={(e) => documentPage(parseInt(e.target.value))} /> of {numPage}</div>
        {barcode ? (
          <button onClick={() => setModal(true)} className="mx-1 px-5 py-2 bg-(--secondary) rounded-xl border border-(--line)">Tandatangani dokumen</button>
        ) : (
          <button onClick={() => setBarcode(true)} className="mx-1 px-5 py-2 bg-(--secondary) rounded-xl border border-(--line)">Tambah barcode</button>
        )}
      </div>
      <div className="w-full overflow-scroll flex justify-center items-center bg-gray-100 border-l border-(--line)">
        <div className="relative scale-75">
          <div style={{width: `${PointToPixel(width)}px`, height: `${PointToPixel(height)}px`, boxShadow: "1px 1px 10px 4px var(--line)"}} ref={pageRef} onClick={handleClick}>
            <Document file={`${path}`} onLoadSuccess={onLoadSuccess}>
                <Page pageNumber={page} width={PointToPixel(width)} />
            </Document>
          </div>
          {barcode && (
            <Sign x={PointToPixel(x)} y={PointToPixel(y)} />
          )}
        </div>
      </div>
      {modal && (
        <Modal >
          <div ref={modalRef} className="bg-background border border-(--line) rounded-2xl">
            <div className="p-5">
              <div className="m-2">
                <p>Masukan passphrase</p>
                <input value={passphrase} onChange={(e) => setPassphrase(e.target.value)} className="border border-(--line) p-2" type="password" />
              </div>
              <button onClick={() => SignDocument()} className="mt-3 mx-1 px-5 py-2 bg-(--secondary) rounded-xl border border-(--line)">Tandatangan</button>
            </div>
          </div>
        </Modal>
      )}
    </section>
  );
}