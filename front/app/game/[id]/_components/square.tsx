import { useMemo } from "react";

type Props = {
	value: number;
	x: number;
	y: number;
	onClick: (row: number, col: number) => void;
};

const Square = (props: Props) => {
	const convertedValue = useMemo(() => {
		if (props.value === 1) {
			return "◯";
		}
		if (props.value === 2) {
			return "✕";
		}
		return "　";
	}, [props.value]);

	return (
		<button
			type="button"
			className={`
      bg-white
      border
      border-4
      border-solid
      border-indigo-500/75
      font-bold
      text-xl
      h-12
      w-12
      -mr-1
      -mt-1
      p-0
      text-center
      `}
			onClick={() => {
				props.onClick(props.x, props.y);
			}}
		>
			{convertedValue}
		</button>
	);
};

export default Square;

// const StyledButton = styled.button`
// background: #fff;
// border: 1px solid #999;
// float: left;
// font-size: 24px;
// font-weight: bold;
// line-height: 34px;
// height: 34px;
// margin-right: -1px;
// margin-top: -1px;
// padding: 0;
// text-align: center;
// width: 34px;
