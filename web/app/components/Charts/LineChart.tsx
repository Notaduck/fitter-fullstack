import { area, bisector, curveCardinal, extent, max, pointer, range, scaleLinear, select, line } from "d3";
import { useContext, useEffect, useMemo, useRef, useState } from "react";
import useMeasure from "react-use-measure";
import { motion } from 'framer-motion'
import { useActivityStore } from "@/store/zustand";

export type Point = {
    x: number;
    y: number;
    id: number
};

type Props = {
    entries: Point[];
    xUnit: string;
    yUnit: string;
    color: string;
    avgValue: number,
    maxValue: number
    xTickFormat?: (t: number) => string;
    yTickFormat?: (t: number) => string;
};

export const AreaChart = ({ entries, ...restProps }: Props) => {
    let [ref, bounds] = useMeasure();

    return (
        <div className="relative h-full w-full" ref={ref}>
            {bounds.width > 0 && (
                <ChartInner
                    data={entries}
                    width={bounds.width}
                    height={bounds.height}
                    {...restProps}
                />
            )}
        </div>
    );
};

type ChartInnerProps = {
    data: Point[];
    avgValue: number,
    maxValue: number,
    width: number;
    height: number;
    xUnit: string;
    yUnit: string;
    color: string;
    xTickFormat?: (t: number) => string;
    yTickFormat?: (t: number) => string;
    fillOpacity?: number;
    marginTop?: number;
    marginRight?: number;
    marginBottom?: number;
    marginLeft?: number;
};

function ChartInner({
    data,
    width,
    height,
    avgValue,
    maxValue,
    xUnit,
    yUnit,
    color,
    xTickFormat = (t) => t.toFixed(1),
    yTickFormat = (t) => t.toString(),
    fillOpacity = 0.8,
    marginTop = 1,
    marginRight = 1,
    marginBottom = 20,
    marginLeft = 40,
}: ChartInnerProps) {

    const { recordId, setRecordId, setRecordValue, recordValue, setRecordPosition, recordPosition } = useActivityStore()



    const xScale = useMemo(
        () =>
            scaleLinear()
                .domain([0, max(data, (d) => d.x) as number])
                .range([marginLeft, width - marginRight]),
        [data, marginLeft, marginRight, width]
    );

    const yScale = useMemo(
        () =>
            scaleLinear()
                .domain(extent(data, (d) => d.y) as [number, number])
                .range([height - marginBottom, marginTop])
                .nice(),
        [data, marginBottom, marginBottom, height]
    );

    const xTicks = useMemo(() => {
        const xMax = xScale.domain()[1];
        let step = 5;
        if (xMax > 99) step = 10;
        if (xMax > 999) step = 20;
        return range(0, xMax, step);
    }, [xScale]);

    const yTicks = useMemo(() => {
        const boundedHeight = height - marginTop - marginBottom;
        const tickPerPixels = 40;
        return yScale.ticks(boundedHeight / tickPerPixels);
    }, [height, marginTop, marginBottom, yScale]);


    const areaPath = useMemo(() => {
        const areaPathGenerator = area<Point>()
            .x((d) => xScale(d.x))
            .y0(yScale.range()[0])
            .y1((d) => yScale(d.y))
            .defined((d) => d.x !== undefined && d.y !== undefined)
            .curve(curveCardinal)
        return areaPathGenerator(data) || "";
    }, [data, xScale, yScale]);

    // Create a ref for the SVG element
    const svgRef = useRef(null);
    const yLineRef = useRef(null)

    const [id, setId] = useState()

    // useEffect to handle mousemove event
    useEffect(() => {
        const svg = select(svgRef.current);

        const listeningRect = svg.append("rect")
            .attr("width", width)
            .attr("height", height)
            .attr("fill", "none")
            .style("pointer-events", "all");

        // Mousemove event handler
        const handleMouseMove = function (event) {
            const [xCoord] = pointer(event, listeningRect.node());
            const bisectDate = bisector(d => d.x).left;
            const x0 = xScale.invert(xCoord);
            const i = bisectDate(data, x0, 1);
            const d0 = data[i - 1];
            const d1 = data[i];
            const d = x0 - d0.x > d1.x - x0 ? d1 : d0;

            const xPos = xScale(d.x);
            const yPos = yScale(d.y);

            // setRecordPosition(xPos)

            // console.log(useActivityStore.getState().recordPosition)


            select(yLineRef.current)
                .attr('x1', xPos ?? useActivityStore.getState().recordPosition)
                .attr('x2', xPos ?? useActivityStore.getState().recordPosition)
                .attr('y1', marginTop)
                .attr('y2', height - marginBottom);


            setRecordId(d.id)
        };

        // Attach mousemove event listener to the listeningRect
        listeningRect.on("mousemove", handleMouseMove);

        // Cleanup: Remove the event listener when the component unmounts
        return () => {
            listeningRect.on("mousemove", null);
        };

    }, [data, width, height, xScale, yScale]);


    const getCurrentYValue = () => {
        const yValue = data.find(e => e.id === recordId)?.y

        if (yValue) return `${yValue.toFixed(2)} ${yUnit}`

        return `00.00 ${yUnit}`
    }



    return (

        <div className="grid grid-cols-10 gap-2">
            <div className="  col-span-1 flex flex-col justify-center align-middle">
                <p>Speed</p>
                <p>avg: {avgValue}</p>
                <p>max: {maxValue}</p>
            </div>
            <div className="  col-span-8">

                <svg
                    ref={svgRef}
                    className="text-[12px]" viewBox={`0 0 ${width} ${height}`}
                >
                    <line ref={yLineRef} className="hover-line" stroke="currentColor" />
                    {data.map((d, i) => (
                        <g
                            id={String(d.id)}
                            key={i}
                            transform={`translate(${xScale(d.x)}, ${yScale(d.y)})`}
                        >
                            {/* Vertical line */}
                            <line
                                id={String(d.id)}
                                className={`hover-line  hover:block  hidden text-purple-600`}
                                x1={d.x}
                                x2={d.x + 1}
                                y1={marginTop}
                                y2={height - marginBottom}
                                stroke="currentColor"
                            />

                        </g>
                    ))}


                    {/* X axis */}
                    <g>
                        {/* X ticks */}
                        {xTicks.map((t, i) => {

                            return (


                                <g key={t} transform={`translate(${xScale(t)},0)`}    >
                                    <line
                                        className="text-gray-300"
                                        y1={marginTop}
                                        y2={height}
                                        stroke="currentColor"
                                    ></line>
                                    {/* Only show text when there's enough space */}
                                    {width - xScale(t) > 48 ? (
                                        <text
                                            className="text-gray-600"
                                            fill="currentColor"
                                            // textAnchor="center"
                                            x="8"
                                            y={height - 4}
                                        >
                                            {i === 0 ? xUnit : xTickFormat(t)}
                                        </text>
                                    ) : null}
                                </g>
                            )
                        })}
                        {/* Last tick */}
                        <line
                            transform={`translate(${xScale.range()[1]},0)`}
                            className="text-gray-300"
                            y1={marginTop}
                            y2={height}
                            stroke="currentColor"
                        ></line>
                    </g>

                    {/* Y axis */}
                    <g>
                        {/* Y ticks */}
                        {yTicks.map((t, i) => (
                            <motion.g
                                key={t}
                                transform={`translate(0,${yScale(t)})`}

                                initial={{ pathLength: 0 }}
                                animate={{ pathLength: 1 }}
                                transition={{
                                    duration: 1.5,
                                    delay: 1,
                                    type: 'spring'
                                }}
                            >
                                <line
                                    className="text-gray-300"
                                    x1={marginLeft}
                                    x2={width - marginRight}
                                    stroke="currentColor"
                                ></line>

                                {height - marginBottom - yScale(t) > 20 &&
                                    yScale(t) - marginTop > 10 ? (
                                    <text
                                        className="text-gray-600"
                                        fill="currentColor"
                                        textAnchor="end"
                                        x={marginLeft - 6}
                                        dy="0.35em"
                                    >
                                        {yTickFormat(t)}
                                    </text>
                                ) : null}
                            </motion.g>
                        ))}
                        {/* First tick */}
                        <g transform={`translate(0,${yScale.range()[0]})`}>
                            <line
                                className="text-gray-300"
                                x2={width - marginRight}
                                stroke="currentColor"
                            ></line>

                            <text
                                className="text-gray-600"
                                fill="currentColor"
                                textAnchor="end"
                                x={marginLeft - 6}
                                y={-6}
                            >
                                {yUnit}
                            </text>
                        </g>
                        {/* Last tick */}
                        <line
                            transform={`translate(0,${yScale.range()[1]})`}
                            className="text-gray-300"
                            x2={width - marginRight}
                            stroke="currentColor"
                            textAnchor="center"
                        ></line>

                        {/* Area */}
                        <motion.path
                            className="" fill={'currentColor'} fillOpacity={fillOpacity} d={areaPath}
                            initial={{ pathLength: 0 }}
                            animate={{ pathLength: 1 }}
                            transition={{
                                duration: 1.5,
                                delay: 1,
                                type: 'spring'
                            }}
                        ></motion.path>

                    </g>
                </svg>
            </div>
            {/* <div className="col-span-1 flex items-center justify-center">{getCurrentYValue()}</div> */}
        </div>
    );
}