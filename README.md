## CoinMarketCap-Cryptocurrency-CLI

### Motivation
- Many Cryptocurrency investors hold Cryptocurrency across multiple investing accounts and hardware wallets. This application performs calculations to determine the total value of their investment. <br>
- This application could perform as an additional service to any preexisting portfolio projects
- Writing a technical blog post about Cobra-cli for my website: https://nicholaspazienza.com/blog <br>

### Example
A user owns .1234 Bitcoin on Coinbase and.25 Bitcoin on their Trezor Hardware Wallet.
The total quantity owned by this user is .3734 Bitcoin.<br>
Once the user knows their quantity they can simply enter the name and quantity into the application to find out their total value. <br>
Command: **go run main.go get --name=Bitcoin --quantity=.3734** <br>

Date & Time: 2025-01-20 00:42:02 <br>
Cryptocurrency Name: Bitcoin <br>
Symbol: BTC <br>
Price: \$102308.01 <br>
Quantity: 0.3734 <br>
Value: \$38201.81 <br> <br>

Summary: <br>
A lot of Cryptocurrencies have become expensive and investors can only purchase fractional shares, it becomes difficult to calculate in your head how much 0.3734 Bitcoin is worth at a price of \$102,308.01. Cryptocurrency is a 24/7 industry with prices constantly changing. As long as a user knows their quantity owned, this CLI tool can use CoinMarketCap's data to provide the value of their investment.  <br> <br>
Additionally, if an investor was interested in purchasing a quarter of a Bitcoin, this application can provide them information about the necessary capital needed to achieve that goal.


### Installation
- Install Go: https://go.dev/doc/install <br>
- Clone the repository: <br>
- Navigate to project directory and install Cobra-cli: <br>
Command: **go install github.com/spf13/cobra-cli@latest**

### How to run the application
- Users will need to get an API key from CoinMarketCap before the project can be ran.
- Use command: **go run main.go get --name=NAME_OF_CRYPTOCURRENCY_HERE --quantity=QUANTITY_OF_CRYPTOCURRENCY** <br>
Ex: **go run main.go get --name=Bitcoin --quantity=.1234** <br>
Ex: **go run main.go get --name=Solana --quantity=32.5**
- Data displayed to console: Name, Symbol, Current USD Price at time of request, Quantity, Value

### Tech Stack
- GoLang (Programming Language)
- GoLand (IDE)
- Cobra-cli: https://github.com/spf13/cobra-cli
- Cobra User Guide: https://github.com/spf13/cobra/blob/main/site/content/user_guide.md#using-the-cobra-library
- CoinMarketCap: https://coinmarketcap.com/
- CoinMarketCap API: https://coinmarketcap.com/api/pricing/
- Git: https://git-scm.com/downloads

### Website
- https://nicholaspazienza.com/

