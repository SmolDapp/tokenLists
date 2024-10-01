# Token List Generators

This document provides an overview of the different generators used to create token and pool lists for our project.

## Generator Types

Our generators are categorized based on their data collection method:

1. **API-based Generators** (GenerationAPI)
2. **External List Generators** (GenerationExternalList)
3. **Event-based Generators** (GenerationEvents)
4. **Chain-specific Generators** (GenerationChain)
5. **Static Generators** (GenerationStatic)
6. **Legacy Generators** (GenerationLegacyList)

## Generator List

### API-based Generators

These generators fetch data by calling external APIs:

- 1inch
- Bebop
- Blockscout
- CoinGecko
- Curve
- Etherscan (excluded due to Cloudflare scraping protection)
- Filecoin
- Messari
- Paraswap
- Portals
- Routescan
- Token Name Service (TNS)
- SushiSwap
- UniSwap
- Yearn
- Yearn Minimal
- ZKSync

### External List Generators

These generators retrieve data from pre-existing lists or external sources:

- Consensys
- Cow Swap
- DefiLlama
- Ledger
- Optimism
- SmolAssets

### Event-based Generators

These generators listen to on-chain events to gather information:

- SushiSwap (token pairs)
- SushiSwap (pools)
- UniSwap (pairs)
- UniSwap (pools)

### OnChain Generators

These generators fetch data directly from on-chain contracts:

- Aerodrome
- Ajna
- Velodrome

### Static Generators

These generators use static, manually curated lists:

- Ajna (Static)
- Filecoin (Static)

## Generator Characteristics

Generators differ in several key aspects:

1. **Generation Method**: API-based, External List, Event-based, Chain-specific, or Static.
2. **Generator Type**: 
   - Token generators (GeneratorToken): Create lists of individual tokens.
   - Pool generators (GeneratorPool): Create lists of liquidity pools or token pairs.
3. **Data Source**: Platform-specific (e.g., 1inch, Curve) or aggregators (e.g., CoinGecko, DefiLlama).
4. **Scope**: Chain-specific (e.g., Optimism, ZKSync) or multi-chain/chain-agnostic.
5. **Purpose**: General token lists or specialized (e.g., Yearn's vault-focused lists).
6. **Curation**: Automatically generated or manually curated ("static" lists).
7. **Inclusion/Exclusion**: Some generators may be marked as excluded (e.g., Etherscan due to Cloudflare scraping protection).

These diverse generators ensure a comprehensive and up-to-date dataset for our project, covering various aspects of the DeFi ecosystem.

## Notes

- Some previously used generators (e.g., ethereum-etherscan, polygon-zkevm) have been deprecated and replaced by more general or updated versions.
- The Etherscan generator is currently excluded due to Cloudflare scraping protection issues.