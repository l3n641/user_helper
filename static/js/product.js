const commonHeaders = {
    'Content-Type': 'application/json;charset=utf-8'
}

async function getProduct(tableName) {
    const url = '/product?table_name=' + tableName;
    try {
        const response = await fetch(url);
        return await response.json();
    } catch (error) {
        console.log('Request Failed', error);
    }
}

async function updateSizeChartBySpu(tableName, spuList, sizeChart) {
    const url = '/size_chart'
    const data = {
        method: 'PUT',
        headers: commonHeaders,
        body: JSON.stringify({spu_list: spuList, size_chart: sizeChart, table_name: tableName})
    }
    try {
        const response = await fetch(url, data);
        return await response.json();
    } catch (error) {
        console.log('Request Failed', error);
    }
}


async function deleteProduct(tableName, emptyField) {
    const url = '/delete_product'
    const data = {
        method: 'POST', headers: commonHeaders, body: JSON.stringify({empty_field: emptyField, table_name: tableName})
    }
    try {
        const response = await fetch(url, data);
        return await response.json();
    } catch (error) {
        console.log('Request Failed', error);
    }
}


async function renameProductCategory(tableName, oldCategory, newCategory) {
    const url = '/rename_category'
    const data = {
        method: 'PUT',
        headers: commonHeaders,
        body: JSON.stringify({old_category: oldCategory, table_name: tableName, new_category: newCategory})
    }
    try {
        const response = await fetch(url, data);
        return await response.json();
    } catch (error) {
        console.log('Request Failed', error);
    }
}
