import { NextResponse } from "next/server";
import { writeFile } from "fs/promises";
import path from "path";

export async function POST(req: Request) {
  const formData = await req.formData();
  const file = formData.get("file") as File;

  if (!file) {
    return NextResponse.json({ error: "File tidak ada" }, { status: 400 });
  }

  const bytes = await file.arrayBuffer();
  const buffer = Buffer.from(bytes);

  const filepath = path.join(process.cwd(), "public/uploads", file.name);

  await writeFile(filepath, buffer);

  return NextResponse.json({
    message: "Upload berhasil",
    filename: file.name,
    url: `/uploads/${file.name}`,
  });
}