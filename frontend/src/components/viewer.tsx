"use client";

import { Document, Page, pdfjs } from "react-pdf";
import "react-pdf/dist/Page/TextLayer.css";
import "react-pdf/dist/Page/AnnotationLayer.css";
import { useState } from "react";
import PointToPixel from "../../util/pointToPixel";

pdfjs.GlobalWorkerOptions.workerSrc = new URL("pdfjs-dist/build/pdf.worker.min.mjs",import.meta.url).toString();

export default function Viewer({ path, pg }: { path: string, pg: number }) {
  const [width, setWidth] = useState<number>(100)
  const [height, setHeight] = useState<number>(100)
  const [numPage, setNumPage] = useState<number>(1)

  const onLoadSuccess = async (pdf: pdfjs.PDFDocumentProxy) => {
    const page = await pdf.getPage(pg)
    const viewport = page.getViewport({ scale: 1 })
    setNumPage(pdf.numPages)

    setWidth(viewport.width)
    setHeight(viewport.height)
  }

  return (
    <div style={{width: `${PointToPixel(width)}px`, height: `${PointToPixel(height)}px`, background: "red"}}>
      <Document file={`${path}`} onLoadSuccess={onLoadSuccess}>
          <Page pageNumber={pg} width={PointToPixel(width)} />
      </Document>
    </div>
  );
}