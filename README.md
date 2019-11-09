# configurator

Prototype/experiment to explore the idea of Ethereum client configurations
as parameterised templates, taking chain spec arguments expressed as a set
of EIPs, which are themselves chain spec arguments applied at certain block
number

Block number 0 EIPs are in this model synonymous with the basic initial 
parameters, like maxCodeSize, gasLimit and so on. 

The idea is that you define a chainspec that references either the initial
parameters of EIP specific values (eg: blockreward and block number in the
following)
```json
"blockReward": {
					"0x0": "0x4563918244f40000",
					"0x42ae50": "0x29a2241af62c0000",
					"0x6f1580": "0x1bc16d674ec80000"
				},
```