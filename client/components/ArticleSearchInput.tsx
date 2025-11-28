"use client";
import { useState, useEffect, useRef } from "react";

interface Article {
    ID: number;
    Title: string;
}

interface ArticleSearchInputProps {
    readonly value: string;
    readonly onChange: (value: string) => void;
    readonly placeholder: string;
}

export default function ArticleSearchInput({
    value,
    onChange,
    placeholder,
}: ArticleSearchInputProps) {
    const [suggestions, setSuggestions] = useState<Article[]>([]);
    const [showSuggestions, setShowSuggestions] = useState(false);
    const [loading, setLoading] = useState(false);
    const containerRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        const fetchSuggestions = async () => {
            if (value.length < 3) {
                setSuggestions([]);
                setShowSuggestions(false);
                return;
            }

            setLoading(true);
            try {
                const res = await fetch(
                    `${process.env.NEXT_PUBLIC_API_URL}/api/articles?title=${encodeURIComponent(value)}&limit=10`
                );
                const data = await res.json();
                if (res.ok) {
                    setSuggestions(data);
                    setShowSuggestions(true);
                }
            } catch (e) {
                console.error("検索候補の取得に失敗しました", e);
            } finally {
                setLoading(false);
            }
        };

        const timeoutId = setTimeout(fetchSuggestions, 300);
        return () => clearTimeout(timeoutId);
    }, [value]);

    useEffect(() => {
        const handleClickOutside = (event: MouseEvent) => {
            if (
                containerRef.current &&
                !containerRef.current.contains(event.target as Node)
            ) {
                setShowSuggestions(false);
            }
        };

        document.addEventListener("mousedown", handleClickOutside);
        return () =>
            document.removeEventListener("mousedown", handleClickOutside);
    }, []);

    const handleSelect = (title: string) => {
        onChange(title);
        setShowSuggestions(false);
        setSuggestions([]);
    };

    return (
        <div ref={containerRef} className="relative w-full">
            <input
                type="text"
                value={value}
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => onChange(e.target.value)}
                placeholder={placeholder}
                required
                className="w-full px-6 py-4 text-black bg-white border-2 border-white rounded-lg outline-none focus:border-[#00ADD8] focus:ring-2 focus:ring-[#00ADD8]/20 transition-all duration-200"
            />
            {showSuggestions && suggestions.length > 0 && (
                <div className="absolute left-0 right-0 z-10 mt-2 bg-white border-2 border-gray-200 rounded-lg max-h-60 overflow-y-auto">
                    {suggestions.map((article) => (
                        <button
                            type="button"
                            key={article.ID}
                            onClick={() => handleSelect(article.Title)}
                            className="w-full text-left px-5 py-3 text-black hover:bg-[#00ADD8] hover:text-white cursor-pointer transition-all duration-150 border-b border-gray-100 last:border-b-0"
                        >
                            {article.Title}
                        </button>
                    ))}
                </div>
            )}
            {loading && showSuggestions && (
                <div className="absolute left-0 right-0 z-10 mt-2 bg-white border-2 border-gray-200 rounded-lg px-5 py-3 text-gray-600 text-sm">
                    読み込み中...
                </div>
            )}
        </div>
    );
}
