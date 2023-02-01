import fs from "fs-extra";
import path from "path";
import Ajv from "ajv";
import addFormats from "ajv-formats";
import ethers from "ethers";
const DataDirectory = "./lists/";
const IndexName = "index.json";
const StandardExtensions = [".json"];

function loadValidators() {
	const ajv = new Ajv.default();
	addFormats(ajv);
	ajv.addFormat("address", (value) => {
		if (ethers.utils.getAddress(value) === value) {
			return true;
		}
		console.error(`Error: "${value}" is not a valid address. Should be ${ethers.utils.getAddress(value)}.`);
		return false
	});
	const validators = {};
	try {
		const schema = JSON.parse(fs.readFileSync("./scripts/schema.json", "utf-8"));
		validators['tokenlistooor'] = ajv.compile(schema);
	} catch (error) {
		console.error(error)
		console.error(`Error: "${schema}" is not a valid schema.`);
		process.exit(1);
	}
	return validators;
}

function validate(directory, validators) {
	let allValid = true;
	for (let name of fs.readdirSync(directory)) {
		if (name.startsWith(".") || name === IndexName) continue;
		const file = path.join(directory, name);
		const ext = path.extname(file);
		const stat = fs.lstatSync(file);

		if (stat.isFile() && StandardExtensions.includes(ext)) {
			try {
				const	data = JSON.parse(fs.readFileSync(file, "utf-8"));
				const	validator = validators['tokenlistooor'];
				if (name.startsWith("0x") && name.endsWith(".json") && name.length === (42 + 5)) {
					const	rawAddress = name.replace(`.json`, ``)
					try {
						if (ethers.utils.getAddress(rawAddress) !== rawAddress) {
							console.error(`Error: "${name}" is not checksummed. ("${file}")`);
							allValid = false;
						}
					} catch {
						console.error(`Error: "${name}" is not a valid address. ("${file}")`);
						allValid = false;
					}
				}

				const valid = validator(data);
				if (!valid) {
					console.error(`Error: "${file}" does not follow schema:`);
					for (const error of validator.errors) {
						if (error.keyword === 'enum') {
							console.log(` - ${error.keyword}: ${error.dataPath} ${error.message}: ${error.params.allowedValues.join(', ')}`);
						} else {
							console.log(` - ${error.keyword}: ${error.dataPath} ${error.message}`);
						}
					}
					allValid = false;
				}
			} catch (error) {
				console.error(`Error: "${file}" is not a valid JSON file: [${error.argument}: ${error.reason} for value ${error.value}]`);
				allValid = false;
				continue;
			}
		} else if (stat.isDirectory()) {
			if (name.startsWith("0x")) {
				try {
					if (ethers.utils.getAddress(name) !== name) {
						console.error(`Error: "${name}" is not checksummed. ("${file}")`);
						allValid = false;
					}
				} catch {
					console.error(`Error: "${name}" is not a valid address. ("${file}")`);
					allValid = false;
				}
			}
			if (name.startsWith("_")) {
				continue
			}
			allValid &= validate(file, validators);
		}
	}
	return allValid;
}

function verify(dataDir) {
	const validators = loadValidators();
	const valid = validate(dataDir, validators);
	if (!valid)
		process.exit(1);
}

const cwd = process.cwd();
if (!fs.existsSync(path.join(cwd, ".git"))) {
	console.error("Error: script should be run in the root of the repo.");
	process.exit(1);
}

try {
	verify(DataDirectory);
	console.log("Ok: all files match schema definitions!");
} catch (error) {
	console.error(error);
	process.exit(1);
}
