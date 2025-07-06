import { CatImage } from "./cat-image"; // 追加
import { fetchImage } from "./fetch-image";

export default async function Home() {
  // APIから画像を取得
  const image = await fetchImage();
  // 画像のURLを渡す
  return <CatImage url={image.url} />;
}