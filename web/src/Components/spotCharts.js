import React from 'react';

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
// import { useEffect } from "react";

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

export const priceChartOptions = {
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
};

export const volumeChartOptions = {
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
};

const priceChartLabels = ["00:00", "01:00", "02:00", "03:00", "04:00", "05:00", "06:00", "07:00", "08:00", "09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00"];
const volumeChartLabels = priceChartLabels

export const priceChartData = {
    labels: priceChartLabels,
    datasets: [{
        lineTension: 0.3,
        label: "actual",
        backgroundColor: 'rgba(93, 156, 236, 0.2)',
        borderColor: '#5D9CEC',
        pointBackgroundColor: '#5D9CEC',
        fill: true,
        borderWidth: 2,
        data: [
            21234.12, 21224.23, 21214.56, 21264.78, 21214.90, 21134.12, 21154.34, 21164.56,
            21174.56, 21184.56, 21214.56, 21224.56, 21234.56, 21244.56, 21264.56, 21284.56,
            21319.56, null, null, null, null, null, null, null,
        ]
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
        data: [
            null, null, null, null, null, null, null, null,
            null, null, null, null, null, null, null, null,
            21319.56, 21344.56, 21374.56, 21420.56, 21515.56, 21624.56, 21744.56, 21850.56,
        ]
    }],
};

export const volumeChartData = {
    labels: volumeChartLabels,
    datasets: [{
        backgroundColor: 'rgba(93, 156, 236, 0.2)',
        borderWidth: 0,
        data: [
            5000, 4000, 6000, 5000, 6000, 4000, 3000, 2000,
            5000, 4000, 6000, 5000, 6000, 7000, 8000, 9000,
            10000, null, null, null, null, null, null, null,
        ]
    }],
};

function SpotCharts(props) {
    // useEffect(() => {
    //     window.initSpotCharts()
    // }, []);

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
                            <a href="#" style={{ marginRight: "12px" }}>1w</a>
                            <a href="#" style={{ marginRight: "12px" }}>1m</a>
                        </h4>
                    </div>
                    <div className="ms-auto">
                        <h1 className="mt-n2 text-end color-sunny-light">{props.forecast}</h1>
                        <h4 className="font-400 text-uppercase mt-n2 font-16 opacity-30 text-end">Forecast</h4>
                    </div>
                </div>
                <div className="chart-container" style={{ width: "100%", height: "200px" }}>
                    <Line options={priceChartOptions} data={priceChartData} />
                </div>
                <div className="chart-container" style={{ width: "100%", height: "100px" }}>
                    <Bar options={volumeChartOptions} data={volumeChartData} />
                </div>
            </div>
        </div>
    )
}

export default SpotCharts