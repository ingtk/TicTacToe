type Props = {
  playerTurn: boolean;
};
const Turn = (props: Props) => {
  const text = props.playerTurn ? "あなたの手番です" : "あいての手番です";
  return (
    <p className="text-center text-lg font-medium text-gray-900">{text}</p>
  );
};
export default Turn;
