package mongo

import (
	"context"
	"fmt"
	"testing"

	"websmee/buyspot/internal/domain"
)

func TestAddAssets(t *testing.T) {
	ctx := context.Background()
	client, _ := Connect(ctx, "mongodb://localhost:27017", "", "")
	assetRepository := NewAssetRepository(client)

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "BTC",
		Order:       1,
		IsAvailable: true,
		Name:        "Bitcoin",
		Description: "Bitcoin (abbreviation: BTC[a] or XBT[b]; sign: ₿) " +
			"is a protocol which implements a highly available, public, permanent, and decentralized ledger. " +
			"In order to add to the ledger, a user must prove they control an entry in the ledger. " +
			"The protocol specifies that the entry indicates an amount of a token, bitcoin with a minuscule b. " +
			"The user can update the ledger, assigning some of their bitcoin to another entry in the ledger. " +
			"Because the token has characteristics of money, it can be thought of as a digital currency.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "ETH",
		Order:       2,
		IsAvailable: true,
		Name:        "Ethereum",
		Description: "Ethereum is a decentralized, open-source blockchain with smart contract functionality. " +
			"Ether (Abbreviation: ETH;[a] sign: Ξ) is the native cryptocurrency of the platform. " +
			"Among cryptocurrencies, ether is second only to bitcoin in market capitalization. " +
			"Ethereum was conceived in 2013 by programmer Vitalik Buterin. " +
			"Additional founders of Ethereum included Gavin Wood, Charles Hoskinson, " +
			"Anthony Di Iorio and Joseph Lubin.[6] In 2014, development work began and was crowdfunded, " +
			"and the network went live on 30 July 2015. " +
			"Ethereum allows anyone to deploy permanent and immutable decentralized applications onto it, " +
			"with which users can interact. Decentralized finance (DeFi) applications provide a broad array " +
			"of financial services without the need for typical financial intermediaries like brokerages, " +
			"exchanges, or banks, such as allowing cryptocurrency users to borrow against their holdings " +
			"or lend them out for interest.[9][10] Ethereum also allows users to create and exchange NFTs, " +
			"which are unique tokens representing ownership of an associated asset or privilege, " +
			"as recognized by any number of institutions. Additionally, " +
			"many other cryptocurrencies utilize the ERC-20 token standard on top of the Ethereum blockchain " +
			"and have utilized the platform for initial coin offerings.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "BNB",
		Order:       3,
		IsAvailable: true,
		Name:        "Binance Coin",
		Description: "Binance coin (BNB) is the exchange token of the Binance crypto exchange. " +
			"It was launched originally on the Ethereum blockchain but later migrated to the Binance Smart Chain, " +
			"now called BNB Chain. Holders of BNB with Binance accounts can access discounted fees on the exchange. " +
			"That means demand for the token is linked to demand for the exchange’s services. " +
			"Therefore, buying BNB can be seen as a bet on the success of the exchange, " +
			"in a similar way to buying a share in a company, " +
			"except that owning the token comes with no ownership rights in the exchange.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "XRP",
		Order:       4,
		IsAvailable: true,
		Name:        "XRP",
		Description: "XRP is the native cryptocurrency of XRP Ledger, which is an open-source, " +
			"public blockchain designed to facilitate faster and cheaper payments. " +
			"Sending payments overseas using the legacy financial system " +
			"typically takes one to four business days and can be expensive. " +
			"If a person uses XRP as a bridging currency, it’s possible to settle cross-border transactions " +
			"in less than five seconds on the open-source XRP Ledger blockchain " +
			"at a fraction of the cost of the more traditional methods.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "ADA",
		Order:       5,
		IsAvailable: true,
		Name:        "Cardano",
		Description: "Launched in 2017, Cardano is billed as a third-generation blockchain, " +
			"following Bitcoin and Ethereum, which were the first- and second-generation blockchains. " +
			"Cardano aims to compete directly with Ethereum and other decentralized application platforms, " +
			"saying that it is a more scalable, secure and efficient alternative. Decentralized applications, " +
			"or dapps, are similar to applications on a smartphone. The main difference is dapps run autonomously " +
			"without a third party operating in the background. They achieve that autonomy by using smart contracts, " +
			"which are computer programs designed specifically to perform a function " +
			"when certain predetermined conditions are met.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "MATIC",
		Order:       6,
		IsAvailable: true,
		Name:        "Polygon",
		Description: "Polygon (MATIC) is the native cryptocurrency that powers the Polygon Network, " +
			"a layer 2 platform created in 2017. Originally called the Matic Network, " +
			"the Polygon Network allows developers to create and deploy their own blockchains " +
			"that are compatible with the Ethereum blockchain with a single click, " +
			"as well as enables other Ethereum-based projects to transfer data and tokens between one another " +
			"using the MATIC sidechain. Think of it as a smaller blockchain that runs in parallel " +
			"with the Ethereum blockchain. The price of Polygon’s token, MATIC, has skyrocketed this year.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "DOGE",
		Order:       7,
		IsAvailable: true,
		Name:        "Dogecoin",
		Description: "Doge is the native cryptocurrency of dogecoin, " +
			"a parody cryptocurrency based on a viral internet meme of a Shiba Inu dog. " +
			"At first, the crypto project was created purely as a mockery of other cryptocurrency projects " +
			"that were being launched at the time.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "SOL",
		Order:       8,
		IsAvailable: true,
		Name:        "Solana",
		Description: "Solana is best known as a competitor to Ethereum, the second-largest blockchain project " +
			"by market capitalization. Like Ethereum, Solana offers a way to build decentralized applications, " +
			"which are similar to normal apps like Twitter and Robinhood, but with the help of blockchains, " +
			"they strip away intermediaries. One of the key problems with Ethereum is " +
			"that it's expensive to execute programs. Ethereum has been building \"layer 2\" technologies " +
			"to get around that problem. Solana aims to fix the scalability issues with what it claims is " +
			"an improved underlying infrastructure that offers faster and cheaper transactions.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "DOT",
		Order:       9,
		IsAvailable: true,
		Name:        "Polkadot",
		Description: "DOT is the native cryptocurrency of Polkadot; a blockchain interoperability " +
			"protocol founded in 2016. It is a sharded blockchain, meaning it connects several " +
			"different chains together in a single network, allowing them to process transactions " +
			"in parallel and exchange data between chains without sacrificing security.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "LTC",
		Order:       10,
		IsAvailable: true,
		Name:        "Litecoin",
		Description: "Litecoin (LTC) is a peer-to-peer cryptocurrency that aims to enable fast and low-cost payments " +
			"to anyone in the world. The Litecoin Project was conceived and created by Charles Lee, " +
			"a former Coinbase employee, with the support of multiple members in the Bitcoin community. " +
			"It was launched on October 13th 2011, and has introduced a number of modifications " +
			"based on the original Bitcoin protocol. The most prominent of these is " +
			"Litecoin’s Proof-of-Work consensus: its algorithm is based on Scrypt, instead of SHA-256d. " +
			"Furthermore, Litecoin has a target block time of 2.5 minutes, and a total supply of 84 million. " +
			"In May 2017, Litecoin adopted Segregated Witness (SegWit). In the same month, " +
			"the first Lightning Network transaction was completed on Litecoin, " +
			"transferring 0.00000001 LTC from Zürich to San Francisco in under one second. " +
			"Future development areas include the addition of privacy-features " +
			"(an implementation of the MimbleWimble protocol), the support of Schnorr Signatures, " +
			"and Taproot (a privacy preserving switchable scripting feature).",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "ICX",
		Order:       11,
		IsAvailable: true,
		Name:        "ICON",
		Description: "ICON (ICX) is a blockchain with the goal of " +
			"\"connecting crypto to the real world and advancing our society towards true hyperconnectivity\". " +
			"ICON aims to reach this goal by connecting independent blockchains and enabling transactions " +
			"between them. To do this, ICON uses the self-developed Loop Fault Tolerance (LFT) consensus mechanism. " +
			"Built on on the Byzantine Fault Tolerant (BFT) Tendermint mechanism, " +
			"LFT improves it by increasing performance through a consolidation of network messages. " +
			"ICON has a total supply of 800,460,000 token. During the ICO, an ERC-20 token, ICX, was sold, " +
			"which was subsequently moved to the mainnet in June 2018. " +
			"ICON considers itself to be governed alike to an indirect democracy. To this end, " +
			"ICON makes use of representation channels, reserve channels and incentives " +
			"to align the votes of nodes and users.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "ARB",
		Order:       12,
		IsAvailable: true,
		Name:        "Arbitrum",
		Description: "Arbitrum is a suite of scaling solutions on Ethereum that utilize Optimistic Rollup. " +
			"It allows users to enjoy faster speed and cheaper transaction costs when interacting with web3 dApps. " +
			"There are 2 Arbitrum chains:\nArbitrum One: Major applications are currently housed on Arbitrum One, " +
			"with a concentration of DeFi protocols.\nArbitrum Nova: Utilizes a milder trust assumption in exchange " +
			"for lower fees. Introduced to cater for high-volume applications such as gaming and social.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "ONT",
		Order:       13,
		IsAvailable: true,
		Name:        "Ontology",
		Description: "Ontology is a high performance public blockchain and distributed collaboration platform. " +
			"Ontology aims at solving the trust problem with blockchains, " +
			"with a prime focus on issues such as identity security and data integrity. " +
			"Ontology's mainnet launch took place in July 2018. " +
			"Some of its core features include a lightweight universal smart contract language (with WASM support), " +
			"the support for multiple encryption algorithms, and the support for multiple consensus algorithms " +
			"(e.g., VBFT, DBFT, RBFT, SBFT, PoW).",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "SXP",
		Order:       14,
		IsAvailable: true,
		Name:        "Swipe",
		Description: "Swipe is a crypto-fiat gateway that enables cryptocurrencies " +
			"to be spent as fiat currencies in real time. Its mission is to " +
			"\"make crypto finance mainstream by connecting existing payment networks to cryptocurrencies\". " +
			"Swipe features a wallet for users to deposit, store, and withdraw their cryptocurrencies. " +
			"Users are able to spend these cryptos with the Swipe Visa debit card at any place " +
			"that supports Visa payments. Swipe also features a utility token, SXP, " +
			"that can be used for paying for transaction fees, as a medium of exchange, " +
			"and to receive discounts on the fees. 80% of transaction and withdrawal fees in SXP token " +
			"are automatically burned on-chain.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "ENJ",
		Order:       15,
		IsAvailable: true,
		Name:        "Enjin Coin",
		Description: "Enjin Coin (ENJ) is an Ethereum-based cryptocurrency that is used " +
			"to directly back the value of next-generation blockchain assets. " +
			"It aims to become \"the gold standard for digital assets\". " +
			"Enjin created the open-source ERC-1155 token standard: " +
			"it enables game editors to convert various types of assets, such as " +
			"currency, real estate, digital art, and gaming items into tokens " +
			"that are usable in blockchain applications. Enjin has built a complete ecosystem " +
			"for on-chain applications, which integrates with its open-source token standard " +
			"and features a unique blockchain explorer (EnjinX). Enjin focuses on adoption for " +
			"non-blockchain participants with various solutions for professionals.",
	}))

	fmt.Println(assetRepository.CreateOrUpdate(ctx, &domain.Asset{
		Symbol:      "TRX",
		Order:       16,
		IsAvailable: true,
		Name:        "TRON",
		Description: "TRON is one of the most widely used public blockchains online, " +
			"with 100 million users and a cumulative total of over 3.4 billion transactions. " +
			"The TRON network aims to decentralize content-sharing and establish a framework " +
			"through which future Web3 platforms will operate. TRON price is updated live on Binance. " +
			"Originally designed as a decentralized content distribution platform, " +
			"TRON has expanded in scope and now functions as one of the biggest decentralized application " +
			"(dApp) blockchains. The TRON ecosystem includes multiple scalability and adoption projects, " +
			"such as its dApp sidechain, Sun Network. The TRON network incorporates a variety of features " +
			"that include dApps, smart contracts, and delegated proof-of-stake (DPoS) consensus, " +
			"as well as stablecoin issuance capabilities. To date, TRON hosts the largest circulating supply " +
			"of USD Tether (USDT) stablecoins, having overtaken the Ethereum network in early 2021.",
	}))
}
