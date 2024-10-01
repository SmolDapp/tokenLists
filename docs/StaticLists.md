# Static Token Lists: Process and Maintenance

Static token lists are predefined collections of token information. Unlike the automated token discovery and management processes used elsewhere in this project, static lists require manual maintenance. This documentation focuses specifically on the management and updating of these static token lists.

## Structure of Static Lists

Static lists typically consist of two main components:

1. A source file (e.g., `generators/static/*-static.go`) containing raw token data.
2. A generated JSON file (e.g., `*-static.json`) that serves as the final output.

## Generation Process

1. **Raw Data Definition**: Token information is initially defined in a Go source file.
2. **Data Processing**: A handler function processes the raw data, often fetching additional information from the blockchain.
3. **JSON Generation**: The processed data is converted into a standardized JSON format.

## Adding a New Token to a Static List

To add a new token to a static list:

1. **Locate the Source File**: Find the appropriate `generators/static/*-static.go` file for your network.

2. **Add Token Information**: Append the new token to the end of the list in the source file. Include:
   - Address (use `common.HexToAddress` for proper formatting)
   - Token URI (if available, or an empty string)

   All other information (name, symbol, decimals, etc.) will be fetched on-chain during the generation process.

3. **Update Timestamp**: If required, update the timestamp in the notice section of the file.

4. **Regenerate the List**: Run the list generation script to update the JSON output file.

5. **Verify**: Check the generated JSON file to ensure the new token appears correctly.

## Best Practices

- Always add new tokens to the end of the list to maintain consistency.
- Ensure addresses are correctly formatted.
- Keep logo URIs up to date and use empty strings if no logo is available.
- Regularly review and update the lists to ensure accuracy.



