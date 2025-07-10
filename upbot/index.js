const url = "https://monkeytype.com";

async function callServer(url) {
  try {
    const response = await fetch(url);
    return response.status;
  } catch (error) {
    console.error("Error calling server:", error);
    return 0; 
  }
}

async function main(url) {

  while (true) {
    const statusCode = await callServer(url);
    if (statusCode === 200) {
      console.log("Server is working");
    } else {
      console.log("Server might be down or unreachable");
    }
    await new Promise((res) => setTimeout(res, 2000)); 
  }
}

main(url);
