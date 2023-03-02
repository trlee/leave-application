class LeaveReport_Helpers{
    // Insert to datatable one row per element of data
    insertRows(id, data){
        const target = document.querySelector('#'+id)
        target.innerHTML = ''
        data.forEach(element => {
            let opt = document.createElement('tr')
            opt.id = element.ID
            opt.innerHTML = `<td>${element.ID}</td>
                             <td>${element.leaveTypeID}</td>
                             <td>${element.entitled}</td>
                             <td>${element.taken}</td>
                             <td>${element.creditExpired}</td>`
            target.appendChild(opt)
        })
    }
    
    // trigger datatable and row click event
    triggerDT(dt, dtBody){
        // trigger datatable
        const table = $('#'+dt).DataTable()
        
    }
}