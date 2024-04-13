"use client";
import { useCallback, useEffect, useMemo, useState } from "react";
import useSWR from "swr";
import Board from "./_components/board";
import useSWRMutation from "swr/mutation";
import GameEnded from "./_components/gameEnd";
import GameResult from "./_components/gameResult";
import Matching from "./_components/matching";
import Turn from "./_components/turn";
import PlayerInfo from "./_components/playerInfo";

type Game = {
  board: number[][];
  playerTurn: boolean;
  currentPlayerId: string;
  hostUserId: string;
  guestUserId: string;
  gameStarted: boolean;
  gameEnded: boolean;
  playerWin: boolean | null;
};

async function fetchGameStatus(url: string) {
  const res = await fetch(`http://localhost:1323/${url}`, {
    credentials: "include",
    method: "GET",
  });
  const data: Game = await res.json();
  return data;
}

async function mutatePlayTurn(
  url: string,
  { arg }: { arg: { y: number; x: number } },
) {
  await fetch(`http://localhost:1323/${url}`, {
    credentials: "include",
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(arg),
  });
  return null;
}

const defaultBoard = [
  [0, 0, 0],
  [0, 0, 0],
  [0, 0, 0],
];
const Game = ({ params }: { params: { id: string } }) => {
  const {
    data,
    board,
    playerTurn,
    currentPlayerId,
    hostUserId,
    guestUserId,
    gameEnded,
    playerWin,
    onClickSquare,
  } = useGame(params.id);

  if (!data) {
    return null;
  }

  if (!data.gameStarted) {
    return <Matching />;
  }

  return (
    <div className="flex flex-row justify-center">
      {/* <div className="mt-8 border border-blue-400"> */}
      <div className="mt-8">
        <div className="mt-4">
          <Turn playerTurn={playerTurn} />
        </div>
        <div className="mt-4">
          <PlayerInfo
            playerTurn={data?.playerTurn}
            currentPlayerId={currentPlayerId}
            hostUserId={hostUserId}
            guestUserId={guestUserId}
          />
        </div>
        <div className="mt-4">
          <Board board={board} onClickSquare={onClickSquare} />
        </div>
        <div className="mt-4">
          <GameResult gameEnded={gameEnded} playerWin={playerWin} />
        </div>
        <div className="mt-4">
          <GameEnded gameEnded={gameEnded} />
        </div>
      </div>
    </div>
  );
};
export default Game;

const useGame = (id: string) => {
  const { data } = useSWR(`game/${id}/status`, fetchGameStatus, {
    refreshInterval: 1000,
  });

  const { trigger } = useSWRMutation(`game/${id}/play_turn`, mutatePlayTurn);

  const currentPlayerId = useMemo(() => {
    if (!data?.currentPlayerId) {
      return null;
    }
    return data?.currentPlayerId;
  }, [data?.currentPlayerId]);

  const [board, setBoard] = useState(defaultBoard);
  const [playerTurn, setPlayerTurn] = useState(false);

  useEffect(() => {
    if (data?.board) {
      setBoard(data?.board);
    }
  }, [data?.board]);

  useEffect(() => {
    setPlayerTurn(data?.playerTurn || false);
  }, [data?.playerTurn]);

  const onClickSquare = useCallback(
    (x: number, y: number) => {
      if (data?.gameEnded) {
        return;
      }
      if (!playerTurn) {
        return;
      }
      if (currentPlayerId === "") {
        return;
      }
      if (!data?.hostUserId || !data?.guestUserId) {
        return;
      }
      if (board[y][x] !== 0) {
        return;
      }
      board[y][x] = currentPlayerId === data?.hostUserId ? 1 : 2;
      setBoard(board);
      setPlayerTurn(false);
      trigger({ y, x });
    },
    [
      board,
      currentPlayerId,
      data?.gameEnded,
      data?.guestUserId,
      data?.hostUserId,
      playerTurn,
      trigger,
    ],
  );

  return {
    data,
    board,
    playerTurn,
    currentPlayerId: data?.currentPlayerId || "",
    hostUserId: data?.hostUserId || "",
    guestUserId: data?.guestUserId || "",
    gameEnded: data?.gameEnded || false,
    playerWin: data?.playerWin === undefined ? null : data?.playerWin,
    onClickSquare,
  };
};
