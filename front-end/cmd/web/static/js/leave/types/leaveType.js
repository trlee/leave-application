const   TypeHelpers = new LeaveType_Helpers(),
        TypeAPI = new LeaveType_API(),
        dt = 'leaveType',
        dtBody = 'leaveTypeBody'

window.addEventListener('DOMContentLoaded', () => {
    // fetch all leave types & update DOM
    TypeAPI.getAllLeaveType().then(resp => {
        console.log(resp)
        TypeHelpers.insertRows(dtBody, resp.data)
        TypeHelpers.triggerDT(dt, dtBody)
    })

    myForm = 'updateLeaveForm'
    const myUpdate = document.querySelector('#updateLeaveSubmit')
    myUpdate.addEventListener('click', () => {
        myData = Common.getForm(myForm)
        console.log('Data:', myData)
        TypeAPI.updateLeaveType(myData).then(resp => {
            console.log(resp)
        })
    })
})

// function showUpdateValues(element) {
//     console.log(element)
function showUpdateValues(ID,leaveType,unpaid,limit,entitlementCalculation,gender,attachmentMandatory,encashmentLeave) {
    console.log("Element JSON: ", ID, leaveType, unpaid, limit, entitlementCalculation, gender, attachmentMandatory, encashmentLeave)
    document.updateLeaveForm.ID.value = ID
    document.updateLeaveForm.name.value = leaveType
    document.updateLeaveForm.isUnpaid.value = unpaid ? 1:0
    document.updateLeaveForm.maxLimit.value = limit
    document.updateLeaveForm.entitlementCalc.value = entitlementCalculation
    document.updateLeaveForm.gender.value = gender
    document.updateLeaveForm.isAttachmentMandatory.value = + attachmentMandatory ? 1:0
    document.updateLeaveForm.isEncashmentLeave.value = + encashmentLeave ? 1:0
    return true
}

function deleteValues(ID) {
    TypeAPI.deleteLeaveType(ID).then(resp => {
        console.log(resp)
    })
    return true
}