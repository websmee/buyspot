import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Navigate } from 'react-router-dom'

import SpotHeader from 'Layouts/spotHeader';
import AssetDescriptionModal from "Components/assetDescriptionModal"
import SpotButtons from "Components/spotButtons"
import SpotCharts from "Components/spotCharts"
import NewsArticle from "Components/newsArticle"
import SpotBuyModal from "Components/spotBuyModal"
import NewsArticleModal from "Components/newsArticleModal"
import ErrorMessage from 'Components/errorMessage';
import { getCurrentSpotsData, getSpotByIndex } from 'Store/reducer';
import Footer from 'Layouts/footer';

function Spot() {
    const dispatch = useDispatch();
    const balance = useSelector((state) => state.balance);
    const spot = useSelector((state) => state.spot);
    const currentSpotsIndex = useSelector((state) => state.currentSpotsIndex);
    const currentSpotsTotal = useSelector((state) => state.currentSpotsTotal);
    const unauthorized = useSelector((state) => state.unauthorized);

    useEffect(() => {
        dispatch(getCurrentSpotsData());
    }, [dispatch]);

    useEffect(() => {
        currentSpotsIndex != 0 && dispatch(getSpotByIndex(currentSpotsIndex));
    }, [currentSpotsIndex]);

    if (unauthorized) {
        return <Navigate to='/login' />
    }

    return (
        <>
            <SpotHeader />

            <div className="page-content header-clear-medium">
                <ErrorMessage />

                {currentSpotsTotal == 0 && <div className="ms-3 me-3 mb-4 alert alert-small shadow-xl bg-fade-gray-dark" role="alert" style={{ borderRadius: "15px" }}>
                    <span style={{ borderRadius: "15px 0 0 15px", left: "0", top: "0", bottom: "0" }}><i className="fa fa-circle-info"></i></span>
                    <strong>No spots found at the moment.</strong>
                </div>}

                {currentSpotsTotal > 0 && <>
                    <SpotCharts
                        assetName={spot.asset.name}
                        assetSymbol={spot.asset.symbol}
                        forecast={spot.priceForecast}
                        chartTimes={spot.chartsDataByQuotes[balance.symbol].times}
                        chartPrices={spot.chartsDataByQuotes[balance.symbol].prices}
                        chartForecast={spot.chartsDataByQuotes[balance.symbol].forecast}
                        chartVolumes={spot.chartsDataByQuotes[balance.symbol].volumes}
                        assetDescriptionModalId="asset-desc-modal"
                    />

                    <SpotButtons
                        activeOrdersCount={spot.activeOrders}
                        assetSymbol={spot.asset.symbol}
                        buyModalId="buy-modal"
                    />

                    {spot.news && spot.news.map((article, i) =>
                        <NewsArticle key={i} modalId={"article-modal-" + i} created={article.created} views={article.views} sentiment={article.sentiment}>
                            {article.title}
                        </NewsArticle>
                    )}
                </>}
            </div>

            <AssetDescriptionModal id="asset-desc-modal" assetName={spot.asset.name} assetSymbol={spot.asset.symbol}>
                {spot.asset.description}
            </AssetDescriptionModal>

            <SpotBuyModal
                id="buy-modal"
                assetName={spot.asset.name}
                assetSymbol={spot.asset.symbol}
                balanceSymbol={balance.symbol}
                amount={spot.buyOrderSettings.amount}
                takeProfit={spot.buyOrderSettings.takeProfit}
                takeProfitOptions={spot.buyOrderSettings.takeProfitOptions}
                stopLoss={spot.buyOrderSettings.stopLoss}
                stopLossOptions={spot.buyOrderSettings.stopLossOptions}
            />

            {spot.news && spot.news.map((article, i) =>
                <NewsArticleModal key={i} id={"article-modal-" + i} created={article.created} views={article.views} title={article.title}>
                    {article.content}
                </NewsArticleModal>
            )}

            <Footer />
        </>
    )
}

export default Spot