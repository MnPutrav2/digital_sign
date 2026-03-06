import { JSX } from "react";

export default function Navside(): JSX.Element {
    return (
        <section className="flex group">
            <div className="w-10 border-r border-(--line) h-screen"></div>
            <div className="w-0 border-(--line) h-screen group-hover:w-40 transition-all duration-150 group-hover:border-r"></div>
        </section>
    )
}