import {
  Text,
  RingProgress,
} from "@mantine/core";

interface RingProgressProps {
  progressValue: number;
}

export default function ProgressCircle({
  progressValue,
}: RingProgressProps) {
  return (
    <RingProgress
      sections={[{ value: calculatePercentage(progressValue), color: getColorByPercentage(progressValue) }]}
      label={
        <Text size="xl" align="center">
          {progressValue}
        </Text>
      }
    ></RingProgress>
  );
}

function calculatePercentage(inputValue: number): number{
    return (100 * inputValue) / 30
}

function getColorByPercentage(inputValue: number): string{
    if(calculatePercentage(inputValue) >= 50){
        return 'green'
    }

    if(calculatePercentage(inputValue) <= 50 && calculatePercentage(inputValue) >= 30){
        return 'yellow'
    }

    return 'red'
}
