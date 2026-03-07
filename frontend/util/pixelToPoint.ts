export default function PixelToPoint(p: number): number {
    const ex = p / 1.333

    return Math.round(ex)
}