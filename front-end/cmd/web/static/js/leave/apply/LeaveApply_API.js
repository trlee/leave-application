const broker = "http://localhost:8881/"

class LeaveApply_API {
    // apply a new leave
    async applyLeave(stringifiedJSON) {
        const url = broker + 'leaveApply';

        const body = {
            method: 'POST',
            body: stringifiedJSON,
        }

        const response = await fetch(url, body)
        const result = await response.json()

        return result
    }

    async submitAttachment(formData) {
        const url = broker + 'leave/uploadLeaveAttachment/';

        const body = {
            method: "POST",
            body: formData,
        }

        const response = await fetch(url, body)
        const result = await response.json()
        return result
    }

    async addLeaveAttachmentToDB(stringifiedJSON) {
        const url = broker + 'insertLeaveAttachment/'

        const body = {
            method: 'POST',
            body: stringifiedJSON,
        }

        const response = await fetch(url, body)
        const result = await response.json()

        return result
    }
}