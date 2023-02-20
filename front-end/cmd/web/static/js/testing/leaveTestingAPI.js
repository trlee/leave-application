class Leave_API {
    async verifyLeave(myEmail, myDuration, myReason) {
        const url = "http://localhost:8081/approveLeave";
        const payload = {
            email: myEmail,
            duration: myDuration,
            reason: myReason
        }

        const body = {
            method: 'POST',
            body: JSON.stringify(payload),
        }

        const response = await fetch(url, body)
        const result = await response.json()

        return result
    }
}