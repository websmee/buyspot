import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import SpotHeader from 'Layouts/spotHeader';
import AssetDescriptionModal from "Components/assetDescriptionModal"
import SpotButtons from "Components/spotButtons"
import SpotCharts from "Components/spotCharts"
import NewsArticle from "Components/newsArticle"
import SpotBuyModal from "Components/spotBuyModal"
import NewsArticleModal from "Components/newsArticleModal"
import ErrorMessage from 'Components/errorMessage';
import { getCurrentSpotsData, getSpotByIndex } from 'Store/reducer';

function Spot() {
    const dispatch = useDispatch();
    const balance = useSelector((state) => state.balance);
    const spot = useSelector((state) => state.spot);
    const currentSpotsIndex = useSelector((state) => state.currentSpotsIndex);

    useEffect(() => {
        dispatch(getCurrentSpotsData());
        currentSpotsIndex > 0 && dispatch(getSpotByIndex(currentSpotsIndex));
    }, [dispatch]);

    return (
        <>
            <SpotHeader />

            {currentSpotsIndex == 0 && <div className="page-content header-clear-medium">
                <div className="ms-3 me-3 mb-4 alert alert-small shadow-xl bg-fade-gray-dark" role="alert" style={{borderRadius: "15px"}}>
                    <span style={{borderRadius: "15px 0 0 15px", left: "0", top: "0", bottom: "0"}}><i className="fa fa-times"></i></span>
                    <strong>No spots found at the moment.</strong>
                </div>
            </div>}

            {currentSpotsIndex > 0 && <div className="page-content header-clear-medium">
                <ErrorMessage />

                <SpotCharts
                    assetName={spot.asset.name}
                    assetTicker={spot.asset.ticker}
                    forecast={spot.priceForecast}
                    chartTimes={spot.chartsData.times}
                    chartPrices={spot.chartsData.prices}
                    chartForecast={spot.chartsData.forecast}
                    chartVolumes={spot.chartsData.volumes}
                    assetDescriptionModalId="asset-desc-modal"
                />

                <SpotButtons
                    activeOrdersCount={spot.activeOrders}
                    assetTicker={spot.asset.ticker}
                    buyModalId="buy-modal"
                />

                {spot.news.map((article, i) =>
                    <NewsArticle key={i} modalId={"article-modal-" + i} created={article.created} views={article.views} sentiment={article.sentiment}>
                        {article.title}
                    </NewsArticle>
                )}
            </div>}

            <AssetDescriptionModal id="asset-desc-modal" assetName={spot.asset.name} assetTicker={spot.asset.ticker}>
                {spot.asset.description}
            </AssetDescriptionModal>

            <SpotBuyModal
                id="buy-modal"
                assetName={spot.asset.name}
                assetTicker={spot.asset.ticker}
                balanceTicker={balance.ticker}
                amount={spot.buyOrderSettings.amount}
                takeProfit={spot.buyOrderSettings.takeProfit}
                takeProfitOptions={spot.buyOrderSettings.takeProfitOptions}
                stopLoss={spot.buyOrderSettings.stopLoss}
                stopLossOptions={spot.buyOrderSettings.stopLossOptions}
            />

            {spot.news.map((article, i) =>
                <NewsArticleModal key={i} id={"article-modal-" + i} created={article.created} views={article.views} title={article.title}>
                    {article.content}
                </NewsArticleModal>
            )}
        </>
    )
}

export default Spot