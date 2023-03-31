package mongo

import (
	"context"
	"fmt"
	"testing"

	"websmee/buyspot/internal/domain"
)

func TestAddAssets(t *testing.T) {
	ctx := context.Background()
	client, _ := Connect(ctx, "mongodb://localhost:27017")
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
}
