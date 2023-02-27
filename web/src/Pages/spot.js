import { useEffect } from 'react';

import AssetDescriptionModal from "Components/assetDescriptionModal"
import SpotButtons from "Components/spotButtons"
import SpotCharts from "Components/spotCharts"
import NewsArticle from "Components/newsArticle"
import SpotBuyModal from "Components/spotBuyModal"
import NewsArticleModal from "Components/newsArticleModal"
import { bind, unbind } from "Services/Utils/stickymobile"

function Spot() {
    useEffect(() => {
        bind();

        return () => {
            unbind();
        };
    }, []);

    return (
        <>
            <div className="page-content header-clear-medium">
                <SpotCharts assetName="Bitcoin" assetTicker="BTC" forecast="+3%" assetDescriptionModalId="asset-desc-modal" />
                <SpotButtons activeOrdersCount="3" assetTicker="BTC" currentSpot="3" spotCount="10" buyModalId="buy-modal" />

                <NewsArticle modalId="article-modal-1" created="1 hour ago" views="15.5k" sentimentIconClass="text-dark fa-face-meh">
                    Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain
                </NewsArticle>
                <NewsArticle modalId="article-modal-2" created="3 hours ago" views="25.5k" sentimentIconClass="text-danger fa-face-frown">
                    Seoul Sanctions North Korea Over Crypto Theft
                </NewsArticle>
                <NewsArticle modalId="article-modal-3" created="5 hours ago" views="12.3k" sentimentIconClass="text-success fa-face-smile">
                    IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed
                </NewsArticle>
                <NewsArticle modalId="article-modal-4" created="7 hours ago" views="34.5k" sentimentIconClass="text-danger fa-face-frown">
                    Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                </NewsArticle>
                <NewsArticle modalId="article-modal-5" created="9 hours ago" views="45.6k" sentimentIconClass="text-success fa-face-smile">
                    India's Central Bank Reveals 50,000 Users and 5,000 Merchants Now Using Digital Rupee
                </NewsArticle>
                <NewsArticle modalId="article-modal-6" created="11 hours ago" views="1.2k" sentimentIconClass="text-dark fa-face-meh">
                    Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                </NewsArticle>
                <NewsArticle modalId="article-modal-7" created="13 hours ago" views="3.4k" sentimentIconClass="text-success fa-face-smile">
                    Rich Dad Poor Dad's Robert Kiyosaki Warns 'Everything Will Crash' — Plans to Buy More Bitcoin
                </NewsArticle>
                <NewsArticle modalId="article-modal-8" created="15 hours ago" views="5.6k" sentimentIconClass="text-success fa-face-smile">
                    Chinese Government Launching National Blockchain Innovation Center
                </NewsArticle>
                <NewsArticle modalId="article-modal-9" created="17 hours ago" views="25.1k" sentimentIconClass="text-danger fa-face-frown">
                    Stiffing the Staker: The SEC's Latest Crackdown on Crypto Innovation
                </NewsArticle>
            </div>

            <AssetDescriptionModal id="asset-desc-modal" assetName="Bitcoin" assetTicker="BTC">
                Bitcoin (abbreviation: BTC[a] or XBT[b]; sign: ₿) is a protocol which implements a highly available,
                public, permanent, and decentralized ledger. In order to add to the ledger, a user must prove they
                control an entry in the ledger. The protocol specifies that the entry indicates an amount of a token,
                bitcoin with a minuscule b. The user can update the ledger, assigning some of their bitcoin to another
                entry in the ledger. Because the token has characteristics of money, it can be thought of as a digital
                currency.
            </AssetDescriptionModal>

            <SpotBuyModal id="buy-modal" assetName="Bitcoin" assetTicker="BTC" balanceTicker="USDT" />

            <NewsArticleModal id="article-modal-1" created="1 hour ago" views="15.5k" title="Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain">
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
                Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain.
            </NewsArticleModal>
            <NewsArticleModal id="article-modal-2" created="3 hour ago" views="25.5k" title="Seoul Sanctions North Korea Over Crypto Theft">
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
                Seoul Sanctions North Korea Over Crypto Theft
            </NewsArticleModal>
            <NewsArticleModal id="article-modal-3" created="5 hour ago" views="12.3k" title="IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed">
                IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed
                IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed
                IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed
                IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed
                IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed
                IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed
                IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed
            </NewsArticleModal>
            <NewsArticleModal id="article-modal-4" created="7 hour ago" views="34.5k" title="Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown">
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
                Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown
            </NewsArticleModal>
            <NewsArticleModal id="article-modal-5" created="9 hour ago" views="45.6k" title="India's Central Bank Reveals 50,000 Users and 5,000 Merchants Now Using Digital Rupee">
                India's Central Bank Reveals 50,000 Users and 5,000 Merchants Now Using Digital Rupee
                India's Central Bank Reveals 50,000 Users and 5,000 Merchants Now Using Digital Rupee
                India's Central Bank Reveals 50,000 Users and 5,000 Merchants Now Using Digital Rupee
                India's Central Bank Reveals 50,000 Users and 5,000 Merchants Now Using Digital Rupee
                India's Central Bank Reveals 50,000 Users and 5,000 Merchants Now Using Digital Rupee
                India's Central Bank Reveals 50,000 Users and 5,000 Merchants Now Using Digital Rupee
            </NewsArticleModal>
            <NewsArticleModal id="article-modal-6" created="11 hour ago" views="1.2k" title="Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes">
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
                Economist Peter Schiff Warns of Financial Crisis and 'Much More Severe Recession' Than the Fed Recognizes
            </NewsArticleModal>
            <NewsArticleModal id="article-modal-7" created="13 hour ago" views="3.4k" title="Rich Dad Poor Dad's Robert Kiyosaki Warns 'Everything Will Crash' — Plans to Buy More Bitcoin">
                Rich Dad Poor Dad's Robert Kiyosaki Warns 'Everything Will Crash' — Plans to Buy More Bitcoin
                Rich Dad Poor Dad's Robert Kiyosaki Warns 'Everything Will Crash' — Plans to Buy More Bitcoin
                Rich Dad Poor Dad's Robert Kiyosaki Warns 'Everything Will Crash' — Plans to Buy More Bitcoin
                Rich Dad Poor Dad's Robert Kiyosaki Warns 'Everything Will Crash' — Plans to Buy More Bitcoin
                Rich Dad Poor Dad's Robert Kiyosaki Warns 'Everything Will Crash' — Plans to Buy More Bitcoin
                Rich Dad Poor Dad's Robert Kiyosaki Warns 'Everything Will Crash' — Plans to Buy More Bitcoin
                Rich Dad Poor Dad's Robert Kiyosaki Warns 'Everything Will Crash' — Plans to Buy More Bitcoin
            </NewsArticleModal>
            <NewsArticleModal id="article-modal-8" created="15 hour ago" views="5.6k" title="Chinese Government Launching National Blockchain Innovation Center">
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
                Chinese Government Launching National Blockchain Innovation Center
            </NewsArticleModal>
            <NewsArticleModal id="article-modal-9" created="17 hour ago" views="25.1k" title="Stiffing the Staker: The SEC's Latest Crackdown on Crypto Innovation">
                Stiffing the Staker: The SEC's Latest Crackdown on Crypto Innovation
                Stiffing the Staker: The SEC's Latest Crackdown on Crypto Innovation
                Stiffing the Staker: The SEC's Latest Crackdown on Crypto Innovation
                Stiffing the Staker: The SEC's Latest Crackdown on Crypto Innovation
                Stiffing the Staker: The SEC's Latest Crackdown on Crypto Innovation
                Stiffing the Staker: The SEC's Latest Crackdown on Crypto Innovation
                Stiffing the Staker: The SEC's Latest Crackdown on Crypto Innovation
            </NewsArticleModal>
        </>
    )
}

export default Spot