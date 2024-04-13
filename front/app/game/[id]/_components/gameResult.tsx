type Props = {
  gameEnded: boolean;
  playerWin: boolean | null;
};

const GameResult = (props: Props) => {
  if (!props.gameEnded) {
    return null;
  }
  if (props.playerWin === null) {
    return null;
  }

  if (props.playerWin) {
    return (
      <p
        className={`
          text-center
          text-lg
          font-medium
          text-sky-500
          animate-bounce
         `}
      >
        あなたの勝ちです
      </p>
    );
  }
  return (
    <p
      className={`
          text-center
          text-lg
          font-medium
          text-gray-500
          animate-pulse
         `}
    >
      あなたの負けです
    </p>
  );
};

export default GameResult;
