import React, { JSX } from "react";
import ReactDOM from "react-dom";

interface Modal {
  children: React.ReactNode
}

export default function Modal({ children }: Modal): JSX.Element {
    return ReactDOM.createPortal(
      <div className="fixed inset-0 backdrop-blur-xs flex justify-center items-center">
          {children}
      </div>,
      document.getElementById("modal-root") as HTMLElement
    );
}