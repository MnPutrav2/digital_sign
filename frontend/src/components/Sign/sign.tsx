import { JSX } from "react"

export default function Sign({x, y}: {x: number, y: number}): JSX.Element {
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