"use client";

import styled from "styled-components";
import Square from "./square";

type GameProps = {
  board: number[][];
  onClickSquare: (x: number, y: number) => void;
};

const Board = (props: GameProps) => {
  const { board, onClickSquare } = props;

  return board.map((row, y) => {
    return (
      <Row key={y}>
        {row.map((v, x) => {
          return (
            <Square
              key={`${x}${y}`}
              x={x}
              y={y}
              value={v}
              onClick={() => {
                onClickSquare(x, y);
              }}
            />
          );
        })}
      </Row>
    );
  });
};

export default Board;

const Row = styled.div`
  &:after {
    clear: both;
    content: "";
    display: table;
  }
`;
