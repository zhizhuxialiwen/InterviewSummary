#include <cstdlib>
#include <cstring>
#include <iostream>
#include <string>
#include "IAPIProviderInterface.h"
#include "IParameterInterface.h"
#include "VariantField.h"

#include "Rank.h"

using namespace seamless;
using namespace std;

namespace seamless {

void Rank::processQuery(IAPIProviderInterface* iAPIProvider) {
        IParameterInterface* iParameter = iAPIProvider->getParameterInterface();
        if (!iParameter) {
                cerr << "iParameter is NULL" << endl;
                return;
        }

        int32_t isRetailWholesale = 0;
        int32_t isUseAlipay = 0;

        VariantField* value = new VariantField;

        iParameter->getParameter("retail_wholesale", value);
        isRetailWholesale = (strcmp(value->strVal, "0")) ? 1 : 0;

        iParameter->getParameter("is_use_alipay", value);
        isUseAlipay = (strcmp(value->strVal, "0")) ? 1 : 0;

        cout << "isRetailWholesale:\t" << isRetailWholesale << endl;
        cout << "isUseAlipay:\t" << isUseAlipay << endl;

        delete value;
        delete iParameter;
}

}  // namespace seamless