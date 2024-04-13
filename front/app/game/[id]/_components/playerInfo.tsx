import { useMemo } from "react";

type Props = {
  playerTurn: boolean;
  currentPlayerId: string;
  hostUserId: string;
  guestUserId: string;
};
const PlayerInfo = (props: Props) => {
  const { me, opponent } = useMemo(() => {
    if (props.playerTurn) {
      if (props.currentPlayerId === props.hostUserId) {
        return { me: "◯", opponent: "✕" };
      }
      if (props.currentPlayerId === props.guestUserId) {
        return { me: "✕", opponent: "◯" };
      }
    } else {
      if (props.currentPlayerId === props.hostUserId) {
        return { me: "✕", opponent: "◯" };
      }
      if (props?.currentPlayerId === props.guestUserId) {
        return { me: "◯", opponent: "✕" };
      }
    }
    return { me: "", opponent: "" };
  }, [props]);
  if (
    props.currentPlayerId === "" ||
    props.hostUserId === "" ||
    props.guestUserId === ""
  ) {
    return null;
  }
  return (
    <>
      <p className="text-center text-lg font-medium text-gray-900">
        あなた：{me}
      </p>
      <p className="text-center text-lg font-medium text-gray-900">
        あいて：{opponent}
      </p>
    </>
  );
};
export default PlayerInfo;
