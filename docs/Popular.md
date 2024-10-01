# Popular Token List

The Popular Token List is a curated collection of widely recognized and frequently used tokens across different blockchain networks. This document explains how the list is created and how tokens are selected for inclusion.

## Purpose

The Popular Token List serves several important purposes:

1. Provide users with quick access to commonly used tokens across various chains.
2. Offer a curated selection of tokens that are recognized and frequently utilized in the ecosystem.
3. Balance token representation across different blockchain networks.

## Token Selection Process

Tokens are chosen for the Popular List by considering their presence across various token lists and adjusting for specific characteristics of each blockchain network.

### 1. Data Aggregation

- Tokens are collected from various source lists (e.g., Uniswap, SushiSwap, CoinGecko, etc.).
- Each token's presence in these lists is recorded.

### 2. List Weighting

- Each source list is assigned a weight based on its size relative to the total number of tokens across all lists.
- This weighting ensures that larger, more comprehensive lists have a proportional influence on token selection.

### 3. Token Frequency Calculation

- For each token, a weighted frequency is calculated.
- When a token appears in a source list, its frequency is incremented by that list's weight.
- Tokens found in multiple lists or in more significant lists will have higher frequencies.

### 4. Chain-Specific Thresholds

- Each blockchain network (chain) has a `WeightRatio` assigned to it.
- This ratio is used to calculate a chain-specific threshold for token inclusion.
- The threshold is computed as: `weightedThreshold = chains.CHAINS[chainID].WeightRatio * sumOfAllListWeights`

### 5. Token Inclusion

- A token is included in the Popular List if its weighted frequency exceeds the chain-specific threshold.
- This method ensures that the list remains relevant across different blockchain ecosystems while accounting for their individual characteristics.

## WeightRatio Explanation

The `WeightRatio` is a crucial factor in determining which tokens are displayed for a given chain:

- **Definition**: The `WeightRatio` is a float value assigned to each chain that adjusts the threshold for including tokens in the popular list.
- **Calculation**: The threshold for token inclusion is calculated as `weightedThreshold = chains.CHAINS[chainID].WeightRatio * sumOfAllListWeights`, where `sumOfAllListWeights` is the sum of weights of all token lists.
- **Impact**:
  - A higher `WeightRatio` results in a higher threshold, making the list more selective and including fewer tokens for that chain.
  - A lower `WeightRatio` leads to a lower threshold, allowing more tokens to be included for that chain.
- **Customization**: This mechanism allows for fine-tuning the Popular List's composition based on the unique aspects of each blockchain network. For example:
  - A chain with numerous tokens might have a higher ratio to maintain a more curated list.
  - A chain with fewer tokens overall might have a lower ratio to ensure adequate representation.


## Benefits of the Popular List

1. **Curation**: Provides a handpicked selection of tokens that are widely recognized and used.
2. **Convenience**: Users can quickly access popular tokens without searching through extensive lists.
3. **Cross-chain Balance**: Ensures fair representation of tokens across different blockchain networks.
4. **Adaptability**: The selection process can be tuned to accommodate changes in the token ecosystem over time.


-------------------

# Chain Lists

In addition to the Popular Token List, individual Chain Lists are also generated and saved. These lists are created using the `SaveChainListInJsonFile` method and provide a detailed collection of tokens specific to each blockchain network. They are using the same list as the Popular list, but with a focus on a specific chain.

### Chain List Creation Process

1. **Token Aggregation**: Tokens are aggregated from various source lists for each specific chain.
2. **Token Validation**: Each token is validated to ensure it meets the criteria for inclusion (e.g., supported chain ID, valid token details).
3. **List Generation**: A list is generated for each chain, containing all valid tokens.
4. **File Saving**: The generated list is saved in a JSON file named after the chain ID (e.g., `1.json` for Ethereum).

### Benefits of Chain Lists

1. **Specificity**: Provides a detailed list of tokens for each blockchain network.
2. **Accuracy**: Ensures that only valid and supported tokens are included.
3. **Convenience**: Users can access a comprehensive list of tokens for a specific chain without sifting through irrelevant data.