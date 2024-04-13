"use client";
import { useRouter } from "next/navigation";
import { useCallback, useState } from "react";
import useSWR from "swr";
import useSWRMutation from "swr/mutation";

type GameStartResponse = {
  game_id: string;
};

async function mutateGameStart() {
  const res = await fetch("http://localhost:1323/game/start", {
    credentials: "include",
    method: "POST",
  });
  const data: GameStartResponse = await res.json();
  return { game_id: data.game_id };
}

export default function Home() {
  const { trigger } = useSWRMutation("/game/start", mutateGameStart);

  const router = useRouter();

  const onClickStartGame = useCallback(async () => {
    const { game_id } = await trigger();
    router.push(`/game/${game_id}`);
  }, [router, trigger]);

  return (
    <>
      <div className="flex flex-row justify-center">
        <div className="mt-8">
          <button
            type="button"
            className={`
            text-white
            bg-gradient-to-br
            from-green-400
            to-blue-600
            active:from-green-500
            active:to-blue-700
            font-medium
            rounded-lg
            text-sm
            px-5
            py-2.5
            text-center
            me-2
            mb-2
            `}
            onClick={onClickStartGame}
          >
            対戦する
          </button>
        </div>
      </div>
    </>
  );
}
