document.getElementById('runBtn').addEventListener('click', runQuery);

async function runQuery() {
    const sql = document.getElementById('sqlInput').value;
    const status = document.getElementById('status');
    const thead = document.getElementById('tableHead');
    const tbody = document.getElementById('tableBody');
    
    thead.innerHTML = '';
    tbody.innerHTML = '';
    status.innerText = "Executing...";

    try {
        const response = await fetch('/api/query', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ sql })
        });

        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(errorText);
        }

        const data = await response.json();
        status.innerText = `Query successful. ${data.length} rows returned.`;

        if (data.length > 0) {
            const columns = Object.keys(data[0]);
            columns.forEach(col => {
                const th = document.createElement('th');
                th.innerText = col;
                thead.appendChild(th);
            });

            data.forEach(row => {
                const tr = document.createElement('tr');
                columns.forEach(col => {
                    const td = document.createElement('td');
                    td.innerText = row[col] === null ? 'NULL' : row[col];
                    tr.appendChild(td);
                });
                tbody.appendChild(tr);
            });
        }
    } catch (err) {
        status.innerText = "Error: " + err.message;
        status.style.color = "#f44747";
    }
}