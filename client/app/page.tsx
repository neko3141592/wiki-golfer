"use client";
import { useState } from "react";
import ArticleSearchInput from "@/components/ArticleSearchInput";

export default function Home() {
    const [start, setStart] = useState("");
    const [end, setEnd] = useState("");
    const [loading, setLoading] = useState(false);
    const [path, setPath] = useState<string[]>([]);
    const [error, setError] = useState("");

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true);
        setError("");
        setPath([]);
        try {
            const res = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/api/path?start=${encodeURIComponent(start)}&end=${encodeURIComponent(end)}`
            );
            const data = await res.json();
            if (!res.ok) {
                setError(data.error || data.message || "エラーが発生しました");
            } else {
                setPath(data.path || []);
            }
        } catch {
            setError("通信エラー");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="min-h-screen bg-zinc-900 text-white flex items-center justify-center p-8 relative">
            {loading && (
                <div className="fixed inset-0 bg-black/70 z-50 flex items-center justify-center">
                    <div className="w-16 h-16 border-4 border-[#00ADD8] border-t-transparent rounded-full animate-spin"></div>
                </div>
            )}
            <div className="container max-w-xl w-full">
                <h1 className="text-5xl font-bold mb-12 tracking-tight text-center">
                    Wiki<span className="text-[#00ADD8] italic">GO</span>lfer
                </h1>
                <form
                    onSubmit={handleSubmit}
                    className="flex flex-col gap-4 mb-12"
                >
                    <ArticleSearchInput
                        value={start}
                        onChange={setStart}
                        placeholder="開始記事タイトルを入力"
                    />
                    <ArticleSearchInput
                        value={end}
                        onChange={setEnd}
                        placeholder="終点記事タイトルを入力"
                    />
                    <button
                        type="submit"
                        disabled={loading}
                        className="px-4 py-4 font-bold bg-white text-black rounded-lg cursor-pointer hover:bg-[#00ADD8] hover:text-white transition-colors duration-300 disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:bg-white disabled:hover:text-black"
                    >
                        {loading ? "検索中..." : "検索"}
                    </button>
                </form>
                <div className="p-4 bg-gradient-to-br from-[#0a0a0a] to-[#1a1a1a] border-2 border-[#333] rounded-xl min-h-[200px] text-base leading-relaxed">
                    {error ? (
                        <div className="flex items-center justify-center h-32">
                            <span className="text-red-400 text-lg">{error}</span>
                        </div>
                    ) : path.length ? (
                        <div className="space-y-2">
                            {path.map((title, index) => (
                                <div key={index}>
                                    <div className="flex items-center gap-4 py-3 px-2 rounded-lg hover:bg-white/5 transition-all duration-200 group">
                                        <span className="text-[#00ADD8] font-bold min-w-[2.5rem] text-right text-lg">
                                            {index + 1}
                                        </span>
                                        <a
                                            href={`https://ja.wikipedia.org/wiki/${encodeURIComponent(
                                                title
                                            )}`}
                                            target="_blank"
                                            rel="noopener noreferrer"
                                            className="text-white  flex-1"
                                        >
                                            {title}
                                        </a>
                                    </div>
                                </div>
                            ))}
                        </div>
                    ) : (
                        <div className="flex items-center justify-center h-32 text-gray-500">
                            結果がここに表示されます
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
}