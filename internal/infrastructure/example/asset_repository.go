package example

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"websmee/buyspot/internal/domain"
)

type AssetRepository struct {
}

func NewAssetRepository() *AssetRepository {
	return &AssetRepository{}
}

func (r *AssetRepository) GetAvailableAssets(ctx context.Context) ([]domain.Asset, error) {
	return []domain.Asset{
		{
			ID:     primitive.NewObjectID(),
			Ticker: "BTC",
			Name:   "Bitcoin",
			Description: "Bitcoin (abbreviation: BTC[a] or XBT[b]; sign: ₿) " +
				"is a protocol which implements a highly available, public, permanent, and decentralized ledger. " +
				"In order to add to the ledger, a user must prove they control an entry in the ledger. " +
				"The protocol specifies that the entry indicates an amount of a token, bitcoin with a minuscule b. " +
				"The user can update the ledger, assigning some of their bitcoin to another entry in the ledger. " +
				"Because the token has characteristics of money, it can be thought of as a digital currency.",
		},
		{
			ID:     primitive.NewObjectID(),
			Ticker: "ETH",
			Name:   "Ethereum",
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
		},
		{
			ID:     primitive.NewObjectID(),
			Ticker: "SHIB",
			Name:   "Shiba Inu token",
			Description: "Shiba Inu token (ticker: SHIB) is a decentralized cryptocurrency created in August 2020 " +
				"by an anonymous person or group known as \"Ryoshi\". It is named after the Shiba Inu (柴犬), " +
				"a Japanese breed of dog originating in the Chūbu region, " +
				"the same breed that is depicted in Dogecoin's symbol, " +
				"itself originally a satirical cryptocurrency based on the Doge meme. " +
				"Shiba Inu has been characterized as a \"meme coin\" and a pump-and-dump scheme. " +
				"There have also been concerns about the concentration of the coin with a single \"whale\" " +
				"wallet controlling billions of dollars' worth of the token, " +
				"and frenzied buying by retail investors motivated by fear of missing out (FOMO).",
		},
	}, nil
}
