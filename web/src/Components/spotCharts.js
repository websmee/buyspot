import React, { useEffect } from 'react';

import {
    BarElement,
    Chart as ChartJS,
    CategoryScale,
    Filler,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { Bar, Line } from 'react-chartjs-2';

import stickymobile from 'Utils/stickymobile';

ChartJS.register(
    CategoryScale,
    Filler,
    LinearScale,
    PointElement,
    LineElement,
    BarElement,
    Title,
    Tooltip,
    Legend
);

function SpotCharts(props) {
    useEffect(() => {
        const menuOpenListener = stickymobile.getMenuOpenListener(props.assetDescriptionModalId);
        const menuCloseListener = stickymobile.getMenuCloseListener();
        stickymobile.bindMenu(props.assetDescriptionModalId, menuOpenListener, menuCloseListener);

        return () => {
            stickymobile.unbindMenu(props.assetDescriptionModalId, menuOpenListener, menuCloseListener);
        }
    }, [])

    let prices = [];
    let forecast = [];
    for(let i = 0; i < props.chartPrices.length; i++) {
        prices.push(props.chartPrices[i]);
        if (i < props.chartPrices.length - 1) {
            forecast.push(null);
        }
    }
    for(let i = 0; i < props.chartForecast.length; i++) {
        forecast.push(props.chartForecast[i]);
        if (i > 0) {
            prices.push(null);
        }
    }

    return (
        <div className="card card-style">
            <div className="content">
                <div className="d-flex">
                    <div>
                        <h1 data-menu={props.assetDescriptionModalId} style={{ cursor: "pointer" }} className="mt-n2">
                            {props.assetName}
                            <span className="font-16 font-400 opacity-50" style={{ marginLeft: "5px" }}>{props.assetTicker}</span>
                        </h1>
                        <h4 className="font-400 text-uppercase mt-n2 font-16 opacity-30">
                            <a style={{ marginRight: "12px" }}>1d</a>
                            <a style={{ marginRight: "12px" }} className="text-info">1w</a>
                            <a style={{ marginRight: "12px" }} className="text-info">1m</a>
                        </h4>
                    </div>
                    <div className="ms-auto">
                        <h1 className="mt-n2 text-end color-sunny-light">+{props.forecast}%</h1>
                        <h4 className="font-400 text-uppercase mt-n2 font-16 opacity-30 text-end">Forecast</h4>
                    </div>
                </div>
                <div className="chart-container" style={{ width: "100%", height: "200px" }}>
                    <Line
                        options={{
                            responsive: true,
                            maintainAspectRatio: false,
                            plugins: {
                                legend: {
                                    display: false,
                                },
                            },
                            title: {
                                display: false
                            },
                            scales: {
                                x: {
                                    display: false
                                }
                            }
                        }}
                        data={{
                            labels: props.chartTimes,
                            datasets: [{
                                lineTension: 0.3,
                                label: "actual",
                                backgroundColor: 'rgba(93, 156, 236, 0.2)',
                                borderColor: '#5D9CEC',
                                pointBackgroundColor: '#5D9CEC',
                                fill: true,
                                borderWidth: 2,
                                data: prices,
                            }, {
                                lineTension: 0.3,
                                label: "forecast",
                                backgroundColor: 'rgba(204, 209, 217, 0.2)',
                                borderColor: 'rgba(204, 209, 217, 0.2)',
                                fill: true,
                                pointStyle: false,
                                pointRadius: 0,
                                borderDash: [5, 5],
                                borderWidth: 2,
                                data: forecast,
                            }],
                        }}
                    />
                </div>
                <div className="chart-container" style={{ width: "100%", height: "100px" }}>
                    <Bar
                        options={{
                            responsive: true,
                            maintainAspectRatio: false,
                            plugins: {
                                legend: {
                                    display: false,
                                },
                            },
                            title: {
                                display: false
                            }
                        }}
                        data={{
                            labels: props.chartTimes,
                            datasets: [{
                                backgroundColor: 'rgba(93, 156, 236, 0.2)',
                                borderWidth: 0,
                                data: props.chartVolumes,
                            }],
                        }}
                    />
                </div>
            </div>
        </div>
    )
}

export default SpotCharts