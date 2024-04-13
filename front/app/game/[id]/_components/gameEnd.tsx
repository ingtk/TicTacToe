type Props = {
  gameEnded: boolean;
};

const GameEnded = (props: Props) => {
  if (!props.gameEnded) {
    return null;
  }

  //   return
  return (
    <>
      <p className="text-center text-lg font-medium">ゲーム終了</p>
      <div className="text-center">
        <a
          href="/"
          className="text-center text-lg font-medium text-blue-500 hover:underline"
        >
          TOPに戻る
        </a>
      </div>
    </>
  );
};

export default GameEnded;
